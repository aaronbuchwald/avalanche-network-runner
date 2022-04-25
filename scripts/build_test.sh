#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# avalanche-network-runner root path
AVALANCHE_NETWORK_RUNNER_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the versions
source "$AVALANCHE_NETWORK_RUNNER_PATH"/scripts/install_avalanchego.sh

go test -timeout="600s" -coverprofile="coverage.out" -covermode="atomic" ./...