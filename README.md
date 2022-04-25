# Avalanche Network Runner

The Avalanche Network Runner provides a simple interface to interact with your own local Avalanche Network on your preferred backend. The Runner has been architected to make it as simple as possible to implement a backend, which can then immediately be used to spin up pre-defined networks.

## Getting Started

Looking to get started using the Avalanche Network Runner?

Right now there are two primary use cases - spin up a network for local development or implement end-to-end tests as simple GoLang unit tests.

Preqrequisites: download and build [AvalancheGo](https://github.com/ava-labs/avalanchego).

```bash
go install github.com/aaronbuchwald/avalanche-network-runner
```

### Create Local Network

To spin up a local network, you can simply run the following:

```bash
avalanche-network-runner
```

or if you want to run with the source code:

```bash
go run main/main.go
```

This will spin up the default five node network on your machine and make the base RPCs of the five nodes available at the following endpoints:

```
http://127.0.0.1:9650
http://127.0.0.1:9652
http://127.0.0.1:9654
http://127.0.0.1:9656
http://127.0.0.1:9658
```

### Create E2E Test

Creating an E2E test using the Avalanche Network Runner is easy and can be done very simply within a GoLang unit test. Currently, these unit tests require that you construct a network orchestrator, spin up a pre-defined or custom network, and defer the teardown of the entire thing to clean up after yourself.

For an example, refer to the existing tests for the network runner itself, which simply check that the default local network becomes healthy on each of the backends ie. [Local Binary Orchestrator Test](./localbinary/orchestrator_test.go).

## Architecture

The Avalanche Network Runner is built on top of the backend `NetworkConstructor` interface. A `NetworkConstructor` creates an "isolated environment" for a group of Avalanche nodes. Isolated is in quotes because it is up to the `NetworkConstructor` to determine how isolated that collection of nodes is.

For example, a network constructor can create all nodes on an isolated Docker Network, so that they will only interact with each other. Or, a network constructor could also be used to spin up nodes on AWS to connect to Fuji/Mainnet.

The `NetworkConstructor` defines how separated this environment actually is. We use as thin of an interface as possible, to keep the task of implementing a new backend as simple as possible.

On top of this, we have the `NetworkBackend`. There is a helper function that takes in a `NetworkConstructor` to form a `NetworkBackend`. The NetworkBackend takes the simple functionality set in `NetworkConstructor` and wraps it to create a more fully featured interface.

This separation ensures that this additional functionality can be easily shared across multiple backends. However, a new backend can also choose to implement its own complete `NetworkBackend` interface, without using this shortcut.

Lastly, there is the `NetworkOrchestrator`, the top-level of the Avalanche Network Runner. The orchestrator simply generates and manages networks. This can be used to create multiple isolated networks simultaneously.

## Backends

The Avalanche Network Runner supports multiple different backends, which define how nodes are created and the environment that they exist in.

Currently, the two possibilities are to start an AvalancheGo binary using the `localbinary` package or to use the Avalanche `Kurtosis Module` to create an isolated Docker Network.

## What's Next for the Avalanche Network Runner?

The Avalanche Network Runner is intended to make it easy build both new backends and new features on top. Here are a couple of future directions that we might take on and would love to see open source contributions on in the meantime.

Avalanche Network Runner is not only a useful tool, it's built to be incredibly easy to contribute to, so this is a great place for new developers to make their first contribution to the Avalanche Ecosystem.

Here are a couple of directions/features it might go in the near future:

1. Build a CLI for interacting with a running network (stay tuned)
2. Build new custom network definitions and tooling to easily create networks with pre-defined genesis and topologies
3. Improving the Kurtosis Module, so that the module itself implements the ability to add a new node to the running network (and leverage Kurtosis SDK to support interacting with the Avalanche Network Runner across different programming languages).
4. Add a new backend to support using the AvalancheGo Application Runner (options: all within the same process, unique process for each node, all running behind a single node RPC server, or each with their own RPC server)
5. Build a new pure Kubernetes Backend to deploy to easily deploy to real infrastructure
