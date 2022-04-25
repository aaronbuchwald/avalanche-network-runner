#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# avalanche-network-runner root path
AVALANCHE_NETWORK_RUNNER_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the versions
source "$AVALANCHE_NETWORK_RUNNER_PATH"/scripts/versions.sh

############################
# download avalanchego
# https://github.com/ava-labs/avalanchego/releases
GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)
DOWNLOAD_URL=https://github.com/ava-labs/avalanchego/releases/download/${avalanchego_version}/avalanchego-linux-${GOARCH}-${avalanchego_version}.tar.gz
DOWNLOAD_PATH=/tmp/avalanchego.tar.gz
if [[ ${GOOS} == "darwin" ]]; then
  DOWNLOAD_URL=https://github.com/ava-labs/avalanchego/releases/download/${avalanchego_version}/avalanchego-macos-${avalanchego_version}.zip
  DOWNLOAD_PATH=/tmp/avalanchego.zip
fi

rm -rf /tmp/avalanchego-${avalanchego_version}
rm -f ${DOWNLOAD_PATH}

echo "downloading avalanchego ${avalanchego_version} at ${DOWNLOAD_URL}"
curl -L ${DOWNLOAD_URL} -o ${DOWNLOAD_PATH}

echo "extracting downloaded avalanchego"
if [[ ${GOOS} == "linux" ]]; then
  tar xzvf ${DOWNLOAD_PATH} -C /tmp
elif [[ ${GOOS} == "darwin" ]]; then
  unzip ${DOWNLOAD_PATH} -d /tmp/avalanchego-build
  mv /tmp/avalanchego-build/build /tmp/avalanchego-${avalanchego_version}
fi
find /tmp/avalanchego-${avalanchego_version}

# Set an environment variable to define the AvalancheGo binary path
export AVALANCHEGO_BINARY_PATH=/tmp/avalanchego-${avalanchego_version}/avalanchego
