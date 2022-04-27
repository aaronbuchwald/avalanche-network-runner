// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import "net"

func GetFreePorts(numPorts int) ([]int, error) {
	ports := make([]int, 0, numPorts)
	for i := 0; i < numPorts; i++ {
		listener, err := net.Listen("tcp", "localhost:0")
		if err != nil {
			return nil, err
		}
		// Defer closing the listener until the end of the function
		// so that we don't accidentally allocate the same port twice.
		defer listener.Close()
		ports = append(ports, listener.Addr().(*net.TCPAddr).Port)
	}

	return ports, nil
}
