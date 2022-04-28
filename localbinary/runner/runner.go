// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package runner

import (
	"context"
	"fmt"
	"sort"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/aaronbuchwald/avalanche-network-runner/backend"
	"github.com/aaronbuchwald/avalanche-network-runner/e2e"
	"github.com/aaronbuchwald/avalanche-network-runner/localbinary"
	"github.com/aaronbuchwald/avalanche-network-runner/networks"
	"github.com/aaronbuchwald/avalanche-network-runner/utils/constants"
	"github.com/aaronbuchwald/avalanche-network-runner/utils/log"
)

// Create and run a five node network
func RunNetwork(ctx context.Context, args []string, networkCallback func(backend.Network) error) error {
	fs := buildFlagSet()
	v, err := buildViper(fs, args)
	if err != nil {
		return err
	}

	level, err := zapcore.ParseLevel(v.GetString(logLevelKey))
	if err != nil {
		return fmt.Errorf("couldn't parse log level: %w", err)
	}
	log.SetGlobalLogLevel(level)

	orchestrator := localbinary.NewNetworkOrchestrator(&localbinary.OrchestratorConfig{
		BaseDir: v.GetString(dataDirectoryKey),
		Registry: map[string]string{
			constants.NormalExecution: v.GetString(avalanchegoBinaryPathKey),
		},
		DestroyOnTeardown: v.GetBool(cleanDataDirKey),
	})
	network, err := networks.NewDefaultLocalNetwork(ctx, orchestrator, constants.NormalExecution)
	if err != nil {
		return err
	}
	defer func() {
		if err := orchestrator.Teardown(context.Background()); err != nil {
			zap.L().Error("Failed to tear down network orchestrator", zap.Error(err))
		}
	}()

	if err := e2e.AwaitHealthy(ctx, network, 5*time.Second); err != nil {
		return err
	}

	zap.L().Info("Network became health...")

	nodes, err := network.GetNodes()
	if err != nil {
		return err
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].GetName() < nodes[j].GetName()
	})
	for _, node := range nodes {
		zap.L().Info("Node became available", zap.String("name", node.GetName()), zap.String("URI", node.GetHTTPBaseURI()))
	}

	// Run the callback on the created network if applicable
	if networkCallback != nil {
		if err := networkCallback(network); err != nil {
			return err
		}
	}

	// Run until the context is marked as done
	<-ctx.Done()
	err = ctx.Err()
	if err == context.Canceled {
		return nil
	}
	return err
}
