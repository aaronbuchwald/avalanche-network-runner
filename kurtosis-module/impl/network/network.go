// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"context"
	"fmt"

	"github.com/aaronbuchwald/avalanche-network-runner/backend"
	"github.com/aaronbuchwald/avalanche-network-runner/kurtosis-module/impl/avalanchego"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/enclaves"
	kurtosiscontext "github.com/kurtosis-tech/kurtosis-engine-api-lib/api/golang/lib/kurtosis_context"
)

var _ backend.NetworkConstructor = &network{}

type network struct {
	kurtosisCtx *kurtosiscontext.KurtosisContext

	enclaveID  enclaves.EnclaveID
	enclaveCtx *enclaves.EnclaveContext

	registry backend.ExecutorRegistry
}

func newNetworkConstructor(
	kurtosisCtx *kurtosiscontext.KurtosisContext,
	enclaveID enclaves.EnclaveID,
	enclaveCtx *enclaves.EnclaveContext,
	registry backend.ExecutorRegistry,
) *network {
	return &network{
		kurtosisCtx: kurtosisCtx,
		enclaveID:   enclaveID,
		enclaveCtx:  enclaveCtx,
		registry:    registry,
	}
}

// AddNode adds a new node to the network defined by [nodeConfig].
func (n *network) AddNode(ctx context.Context, nodeConfig backend.NodeConfig) (backend.Node, error) {
	executable, ok := n.registry.GetExecutor(nodeConfig.Executable)
	if !ok {
		return nil, fmt.Errorf("failed to get executor under name %s", nodeConfig.Executable)
	}
	nodeConfig.Executable = executable
	return avalanchego.LaunchAvalancheGo(n.enclaveCtx, nodeConfig) // Does this take in a ctx?
}

// Teardown destroys all of the resources associated with this network
func (n *network) Teardown(ctx context.Context) error {
	return n.kurtosisCtx.DestroyEnclave(ctx, n.enclaveID)
}
