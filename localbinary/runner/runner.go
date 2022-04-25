// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package runner

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/aaronbuchwald/avalanche-network-runner/backend"
	"github.com/aaronbuchwald/avalanche-network-runner/e2e"
	"github.com/aaronbuchwald/avalanche-network-runner/localbinary"
	"github.com/aaronbuchwald/avalanche-network-runner/networks"
	"github.com/aaronbuchwald/avalanche-network-runner/utils/constants"
)

// Create and run a five node network
func RunNetwork(ctx context.Context, args []string) error {
	fs := buildFlagSet()
	v, err := buildViper(fs, args)
	if err != nil {
		return err
	}

	level, err := logrus.ParseLevel(v.GetString(logLevelKey))
	if err != nil {
		return fmt.Errorf("couldn't parse log level: %w", err)
	}
	logrus.SetLevel(level)

	registry := backend.NewExecutorRegistry(map[string]string{
		constants.NormalExecution: v.GetString(avalanchegoBinaryPathKey),
	})

	orchestrator := localbinary.NewNetworkOrchestrator(v.GetString(dataDirectoryKey), registry, v.GetBool(cleanDataDirKey))
	network, err := networks.NewDefaultLocalNetwork(ctx, orchestrator, constants.NormalExecution)
	if err != nil {
		return err
	}
	defer func() {
		if err := orchestrator.Teardown(context.Background()); err != nil {
			logrus.Errorf("Failed to tear down network orchestrator due to %s\n", err)
		}
	}()

	if err := e2e.AwaitHealthy(ctx, network, 5*time.Second); err != nil {
		return err
	}

	logrus.Info("Network became healthy...\n")

	nodes := network.GetNodes()
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].GetName() < nodes[j].GetName()
	})
	for _, node := range nodes {
		logrus.Infof("%s available at %s.", node.GetName(), node.GetHTTPBaseURI())
	}

	// Run until the context is marked as done
	<-ctx.Done()
	err = ctx.Err()
	if err == context.Canceled {
		return nil
	}
	return err
}