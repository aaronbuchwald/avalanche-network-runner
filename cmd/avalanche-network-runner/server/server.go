// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aaronbuchwald/avalanche-network-runner/grpc/server"
	"github.com/aaronbuchwald/avalanche-network-runner/localbinary"
	"github.com/aaronbuchwald/avalanche-network-runner/utils/constants"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cobra.EnablePrefixMatching = true
}

var (
	logLevel              string
	port                  string
	gwPort                string
	dialTimeout           time.Duration
	orchestratorBaseDir   string
	teardownOnExit        bool
	avalancheGoBinaryPath string
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server [options]",
		Short: "Start a network runner server.",
		RunE:  serverFunc,
	}

	cmd.PersistentFlags().StringVar(&logLevel, "log-level", logrus.InfoLevel.String(), "Log level to use for the server.")
	cmd.PersistentFlags().StringVar(&port, "port", ":8080", "server port")
	cmd.PersistentFlags().StringVar(&gwPort, "grpc-gateway-port", ":8081", "grpc-gateway server port")
	cmd.PersistentFlags().DurationVar(&dialTimeout, "dial-timeout", 10*time.Second, "server dial timeout")
	cmd.PersistentFlags().StringVar(&orchestratorBaseDir, "base-directory", constants.BaseDataDir, "Set the base directory for the orchestrator running behind the server.")
	cmd.PersistentFlags().BoolVar(&teardownOnExit, "destroy-on-teardown", false, "Set boolean on whether or not all data associated with the orchestrator should be destroyed on shutdown.")
	cmd.PersistentFlags().StringVar(&avalancheGoBinaryPath, "avalanchego-binary-path", constants.AvalancheGoBinary, "Sets the path to use for the AvalancheGo binary.")

	return cmd
}

func serverFunc(cmd *cobra.Command, args []string) (err error) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)

	orchestrator := localbinary.NewNetworkOrchestrator(&localbinary.OrchestratorConfig{
		BaseDir: orchestratorBaseDir,
		Registry: map[string]string{
			constants.NormalExecution: avalancheGoBinaryPath,
		},
		DestroyOnTeardown: teardownOnExit,
	})

	s, err := server.New(server.Config{
		Port:        port,
		GwPort:      gwPort,
		DialTimeout: dialTimeout,
	}, orchestrator)
	if err != nil {
		return err
	}

	rootCtx, rootCancel := context.WithCancel(context.Background())
	errc := make(chan error)
	go func() {
		errc <- s.Run(rootCtx)
	}()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-sigc:
		logrus.Warnf("signal %v received; closing server", sig)
		rootCancel()
		err := <-errc
		if err != nil {
			logrus.Warnf("Error while closing server: %s", err)
		} else {
			logrus.Info("Closed server.")
		}
	case err = <-errc:
		logrus.Warn("Closed server with error: %s", err)
		rootCancel()
	}
	return err
}
