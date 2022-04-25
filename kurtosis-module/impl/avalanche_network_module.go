// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package impl

import (
	"time"

	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/enclaves"
	"github.com/sirupsen/logrus"
)

type AvalancheNetworkModule struct{}

func NewAvalancheNetworkModule() *AvalancheNetworkModule { return &AvalancheNetworkModule{} }

func (e *AvalancheNetworkModule) Execute(enclaveCtx *enclaves.EnclaveContext, serializedParams string) (serializedResult string, resultError error) {
	logrus.Infof("AvalancheNetworkModule received serialized execute params %v at %v", serializedParams, time.Now())
	return "", nil
}
