// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanchego

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aaronbuchwald/avalanche-network-runner/backend"
	"github.com/ava-labs/avalanchego/config"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/enclaves"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/services"
	"github.com/sirupsen/logrus"
)

const (
	// TODO: replace PortIDs with flags from AvalancheGo when Kurtosis removes the requirement that the ID cannot contain a '-' character
	// TODO: Remove the default HTTP and Staking Ports when added into AvalancheGo
	HTTPPortID             = "HTTPPortID"
	defaultHTTPPort    int = 9650
	StakingPortID          = "StakingPortID"
	defaultStakingPort int = 9651
)

var _ backend.Node = &kurtosisAvalancheNode{}

type kurtosisAvalancheNode struct {
	enclaveCtx  *enclaves.EnclaveContext
	serviceCtx  *services.ServiceContext
	name        string
	config      map[string]interface{}
	baseURI     string
	bootstrapIP string
}

func (node *kurtosisAvalancheNode) GetName() string { return node.name }

func (node *kurtosisAvalancheNode) GetHTTPBaseURI() string { return node.baseURI }

func (node *kurtosisAvalancheNode) GetBootstrapIP() string { return node.bootstrapIP }

func (node *kurtosisAvalancheNode) Config() map[string]interface{} {
	return backend.CopyConfig(node.config)
}

func (node *kurtosisAvalancheNode) Stop(stopTimeout time.Duration) error {
	return node.enclaveCtx.RemoveService(node.serviceCtx.GetServiceID(), uint64(stopTimeout))
}

func getContainerConfigSupplier(nodeDef backend.NodeConfig) func(ipAddr string, sharedDirectory *services.SharedPath) (*services.ContainerConfig, error) {
	return func(ipAddr string, sharedDirectory *services.SharedPath) (*services.ContainerConfig, error) {
		// Create a new container config
		configBuilder := services.NewContainerConfigBuilder(nodeDef.Executable)

		httpPort := defaultHTTPPort
		if val, ok := nodeDef.Config[config.HTTPPortKey]; ok {
			httpPort = val.(int)
		}
		stakingPort := defaultStakingPort
		if val, ok := nodeDef.Config[config.StakingPortKey]; ok {
			stakingPort = val.(int)
		}

		// Bind the staking and http ports for the Docker container
		configBuilder.WithUsedPorts(map[string]*services.PortSpec{
			HTTPPortID:    services.NewPortSpec(uint16(httpPort), services.PortProtocol_TCP),
			StakingPortID: services.NewPortSpec(uint16(stakingPort), services.PortProtocol_TCP),
		})

		nodeConfigBytes, err := json.Marshal(nodeDef.Config)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal node config: %w", err)
		}

		logrus.Infof("Starting node with image %s at IP Address %s: \n%v\n", nodeDef.Executable, ipAddr, string(nodeConfigBytes))

		// Set to use local network and staking-enabled false to prevent laptop from spinning during testing
		cmdParams := []string{
			"./avalanchego",
			fmt.Sprintf("--%s=%s", config.ConfigContentKey, base64.StdEncoding.EncodeToString(nodeConfigBytes)),
			fmt.Sprintf("--%s=json", config.ConfigContentTypeKey),
		}
		configBuilder.WithCmdOverride(cmdParams)

		return configBuilder.Build(), nil
	}
}

// LauncheAvalancheGo will start a node in [enclaveCtx] with [nodeDef] as the definition and using [bootNodes] as the specified bootstrappers.
func LaunchAvalancheGo(enclaveCtx *enclaves.EnclaveContext, nodeDef backend.NodeConfig) (backend.Node, error) {
	start := time.Now()
	defer func() { logrus.Infof("Launching Node %s took %v", nodeDef.Name, time.Since(start)) }()

	serviceCtx, err := enclaveCtx.AddService(services.ServiceID(nodeDef.Name), getContainerConfigSupplier(nodeDef))
	if err != nil {
		return nil, err
	}

	httpPort, ok := nodeDef.Config[config.HTTPPortKey]
	if !ok {
		httpPort = defaultHTTPPort
	}
	baseURI := fmt.Sprintf("http://%s:%d", serviceCtx.GetServiceID(), httpPort)

	stakingPort, ok := nodeDef.Config[config.StakingPortKey]
	if !ok {
		stakingPort = defaultStakingPort
	}
	bootstrapIP := fmt.Sprintf("%s:%d", serviceCtx.GetPrivateIPAddress(), stakingPort)

	return &kurtosisAvalancheNode{
		enclaveCtx:  enclaveCtx,
		serviceCtx:  serviceCtx,
		name:        nodeDef.Name,
		config:      nodeDef.Config,
		baseURI:     baseURI,
		bootstrapIP: bootstrapIP,
	}, nil
}
