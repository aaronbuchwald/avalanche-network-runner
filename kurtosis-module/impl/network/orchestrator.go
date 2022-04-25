// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aaronbuchwald/avalanche-network-runner/backend"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/enclaves"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/modules"
	kurtosiscontext "github.com/kurtosis-tech/kurtosis-engine-api-lib/api/golang/lib/kurtosis_context"
	"github.com/sirupsen/logrus"
)

var _ backend.NetworkOrchestrator = &networkOrchestrator{}

type networkOrchestrator struct {
	lock sync.RWMutex

	registry    backend.ExecutorRegistry
	kurtosisCtx *kurtosiscontext.KurtosisContext

	moduleID     modules.ModuleID
	moduleImage  string
	moduleParams string

	networks map[string]backend.Network
}

// NewNetworkOrchestrator creates a new network orchestrator relying on the kurtosis engine
func NewNetworkOrchestrator(
	registry backend.ExecutorRegistry,
	moduleID modules.ModuleID,
	moduleImage string,
	moduleParams string,
) (backend.NetworkOrchestrator, error) {
	kurtosisCtx, err := kurtosiscontext.NewKurtosisContextFromLocalEngine()
	if err != nil {
		return nil, err
	}

	return &networkOrchestrator{
		registry:     registry,
		kurtosisCtx:  kurtosisCtx,
		moduleID:     moduleID,
		moduleImage:  moduleImage,
		moduleParams: moduleParams,
		networks:     make(map[string]backend.Network),
	}, nil
}

// CreateNetwork constructs an isolated network given by [name]
func (net *networkOrchestrator) CreateNetwork(name string) (networkBackend backend.Network, err error) {
	net.lock.Lock()
	defer net.lock.Unlock()

	ctx := context.Background()

	if _, ok := net.networks[name]; ok {
		return nil, fmt.Errorf("cannot create duplicate network under the same name: %s", name)
	}

	enclaveID := enclaves.EnclaveID(name)
	enclaveCtx, err := net.kurtosisCtx.CreateEnclave(ctx, enclaveID, false)
	if err != nil {
		return nil, err
	}
	// Defer destroying the enclave if this function exists with a non-nil error to avoid a leak
	defer func() {
		if err != nil {
			if err := net.kurtosisCtx.DestroyEnclave(ctx, enclaveID); err != nil {
				logrus.Errorf("Failed to destroy enclave after failing to create network: %s", err)
			}
		}
	}()

	logrus.Infof("Network orchestrator (EnclaveID: %s) loading module (ID: %s, Image: %s, Params: %v).", enclaveID, net.moduleID, net.moduleImage, net.moduleParams)
	moduleCtx, err := enclaveCtx.LoadModule(net.moduleID, net.moduleImage, net.moduleParams)
	if err != nil {
		return nil, err
	}

	response, err := moduleCtx.Execute("{}")
	if err != nil {
		return nil, err
	}
	logrus.Infof("Executed module to create Kurtosis Avalanche Network Constructor: %v", response)

	networkConstructor := newNetworkConstructor(net.kurtosisCtx, enclaveID, enclaveCtx, net.registry)
	network := backend.NewNetwork(networkConstructor)
	net.networks[name] = network
	return network, nil
}

// GetNetwork returns the network associated with [name]
func (net *networkOrchestrator) GetNetwork(name string) (backend.Network, bool) {
	net.lock.RLock()
	defer net.lock.RUnlock()

	network, exists := net.networks[name]
	return network, exists
}

// Teardown destroys all of the networks that have been created by the orchestrator
func (net *networkOrchestrator) Teardown(ctx context.Context) error {
	start := time.Now()
	defer func() { logrus.Infof("Took %v to teardown network", time.Since(start)) }()
	// TODO: clean should return type map[ID]struct{} instead of bool
	_, err := net.kurtosisCtx.Clean(ctx, true)
	return err
}
