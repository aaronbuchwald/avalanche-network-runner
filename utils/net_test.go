// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFreePorts(t *testing.T) {
	numPorts := 2
	ports, err := GetFreePorts(numPorts)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, ports, numPorts)

	for _, port := range ports {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			t.Fatal(err)
		}
		if err := listener.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
