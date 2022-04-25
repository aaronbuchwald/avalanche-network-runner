// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"fmt"
	"os"

	"github.com/aaronbuchwald/avalanche-network-runner/kurtosis-module/impl"
	"github.com/kurtosis-tech/kurtosis-module-api-lib/golang/lib/execution"
	"github.com/sirupsen/logrus"
)

const (
	successExitCode = 0
	failureExitCode = 1
)

func main() {
	// >>>>>>>>>>>>>>>>>>> REPLACE WITH YOUR OWN CONFIGURATOR <<<<<<<<<<<<<<<<<<<<<<<<
	configurator := impl.NewNewKurtosisAvalancheModuleConfigurator()
	// >>>>>>>>>>>>>>>>>>> REPLACE WITH YOUR OWN CONFIGURATOR <<<<<<<<<<<<<<<<<<<<<<<<

	executor := execution.NewKurtosisModuleExecutor(configurator)
	if err := executor.Run(); err != nil {
		logrus.Errorf("An error occurred running the Kurtosis module executor:")
		fmt.Fprintln(logrus.StandardLogger().Out, err)
		os.Exit(failureExitCode)
	}
	os.Exit(successExitCode)
}
