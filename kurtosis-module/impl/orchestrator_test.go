// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package impl

import (
	"context"
	"testing"
	"time"

	"github.com/aaronbuchwald/avalanche-network-runner/backend"
	"github.com/aaronbuchwald/avalanche-network-runner/e2e"
	"github.com/aaronbuchwald/avalanche-network-runner/kurtosis-module/impl/network"
	"github.com/aaronbuchwald/avalanche-network-runner/utils/constants"
	"github.com/sirupsen/logrus"
)

func TestKurtosisNetworkOrchestrator(t *testing.T) {
	t.Skip("Skipping until docker container teardown has been parallelized (next Kurtosis engine release)")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Minute))
	defer cancel()

	registry := backend.NewExecutorRegistry(map[string]string{
		constants.NormalExecution: constants.AvalancheGoDockerImage,
	})
	logrus.Info("Starting Kurtosis Network Orchestrator")
	orchestrator, err := network.NewNetworkOrchestrator(registry, "avalanche-module", "avaplatform/avalanche-module", "{}")
	if err != nil {
		t.Fatal(err)
	}

	e2e.TestNetworkOrchestrator(ctx, t, orchestrator)
}
