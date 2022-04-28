// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package client

import (
	"context"
	"sync"
	"time"

	"github.com/aaronbuchwald/avalanche-network-runner/backend"
	"github.com/aaronbuchwald/avalanche-network-runner/rpcpb"
	"github.com/aaronbuchwald/avalanche-network-runner/utils/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	LogLevel    string
	Endpoint    string
	DialTimeout time.Duration
}

type Client interface {
	backend.NetworkOrchestrator
	Ping(ctx context.Context) (*rpcpb.PingResponse, error)
	Close() error
}

type client struct {
	cfg Config

	conn *grpc.ClientConn

	pingc rpcpb.PingServiceClient

	orchestratorc rpcpb.OrchestratorServiceClient
	backend.NetworkOrchestrator

	closed    chan struct{}
	closeOnce sync.Once
}

func New(cfg Config) (Client, error) {
	level, err := zapcore.ParseLevel(cfg.LogLevel)
	if err != nil {
		return nil, err
	}
	log.SetGlobalLogLevel(level)

	zap.L().Info("Dialing grpc server", zap.String("endpoint", cfg.Endpoint))

	ctx, cancel := context.WithTimeout(context.Background(), cfg.DialTimeout)
	conn, err := grpc.DialContext(
		ctx,
		cfg.Endpoint,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	cancel()
	if err != nil {
		return nil, err
	}

	orchestratorc := rpcpb.NewOrchestratorServiceClient(conn)
	orchestratorBackend := newOrchestrator(orchestratorc)
	orchestrator := backend.NewOrchestrator(orchestratorBackend)

	return &client{
		cfg:                 cfg,
		conn:                conn,
		pingc:               rpcpb.NewPingServiceClient(conn),
		orchestratorc:       orchestratorc,
		NetworkOrchestrator: orchestrator,
		closed:              make(chan struct{}),
	}, nil
}

func (c *client) Ping(ctx context.Context) (*rpcpb.PingResponse, error) {
	zap.L().Info("Sending ping...")

	// ref. https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/
	// curl -X POST -k http://localhost:8081/v1/ping -d ''
	return c.pingc.Ping(ctx, &rpcpb.PingRequest{})
}

func (c *client) Close() error {
	c.closeOnce.Do(func() {
		close(c.closed)
	})
	return c.conn.Close()
}
