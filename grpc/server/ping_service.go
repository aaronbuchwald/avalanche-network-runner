// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package server

import (
	"context"
	"os"

	"github.com/aaronbuchwald/avalanche-network-runner/rpcpb"
	"github.com/sirupsen/logrus"
)

type PingServiceHandler struct {
	rpcpb.UnimplementedPingServiceServer
}

func (h *PingServiceHandler) Ping(ctx context.Context, in *rpcpb.PingRequest) (*rpcpb.PingResponse, error) {
	logrus.Debugf("Received ping message.")
	return &rpcpb.PingResponse{Pid: int32(os.Getpid())}, nil
}
