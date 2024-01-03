# Building the BuildBeaver CLI

The BuildBeaver CLI can be built on Linux, Mac or Windows.

## Pre-requisites:

If building on Windows, install:

1. **gcc**: Install the GCC compiler. This can be installed as part of mingw e.g. `choco install mingw --version 5.4.0`.
   If using Go 1.19 on Windows then an older version of gcc is required; version 5.4.0 recommended.
   (See Golang [Issue 35006]( https://github.com/golang/go/issues/35006) and [Issue 51007](https://github.com/golang/go/issues/51007)
   for details; this was fixed in Go 1.20.)

2. **make**: A makefile is used for the initial build from source, e.g `choco install make`

For all platforms (including Windows), install:

3. **Go 1.19** or later: 1.19 recommended but newer versions should work.
   See [Go Download and Install](https://go.dev/doc/install). 
 
4. **Docker**: recommended for testing and required for using bb to build itself.
   See [Get Docker](https://docs.docker.com/get-docker/).

5. **Wire**: required for code generation. Type `go install github.com/google/wire/cmd/wire@latest`

## Building the 'bb' command-line tool

1. `cd backend`
1. `make generate`
1. `make build`

Then type 'bb' to check that the command-line tool runs; this should output the help text.

## Building using bb

Once you have built or downloaded the bb executable, it can be used to build itself.

To build just the bb executable, ensure Docker is running and then in the root of the repo type `bb run build`

To run the entire build, including OpenAPI client code generation, type `bb run`

## Running Go Tests

To run the Go tests without using the bb command-line tool:

1. `cd backend`

2. To run all tests: `go test -mod=vendor ./...`.

3. To run unit tests only: `go test -mod=vendor -short ./...`

4. To run integration tests only: `go test -mod=vendor -run Integration ./...`

Note that the bb-cli repo does not include most of the integration tests since they require the full
BuildBeaver CI server.

## OpenAPI Code Generation

The Go Dynamic SDK is in the [go-sdk repo](https://github.com/buildbeaver/go-sdk) and includes OpenAPI-generated code for the low-level client.
The source YAML for this API (dynamic-openapi.yaml) is in this repo.
The docker-based version of *openapi-generator* is used for generating the OpenAPI client code, so there's no need to
install the tool natively.

To regenerate the Go SDK code after changing the dynamic-openapi.yaml file, type: `bb run openapi`.
This will generate code in the `sdk/go-sdk/bb/client` directory.
