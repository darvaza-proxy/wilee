# Protocol Buffers

This directory contains protocol buffer definitions for the wilee project.

## Module

Published as [buf.build/darvaza/wilee](https://buf.build/darvaza/wilee).

## Structure

```text
proto/
├── buf.yaml           # buf module configuration
├── buf.gen-go.yaml    # Go generation template
├── Makefile
└── acme/v1/           # ACME protocol (RFC 8555)
    ├── acme.proto     # Identifier, Directory
    ├── problem.proto  # ErrorType, Problem, Subproblem
    ├── account.proto
    ├── order.proto
    ├── authorization.proto
    └── challenge.proto
```

## Usage

### Go Code Generation

```sh
make generate-go OUT=../pkg/specs
```

This generates Go code to the specified output directory. The template
`buf.gen-go.yaml` configures the `buf.build/protocolbuffers/go` plugin.

### Linting

```sh
make lint
```

Or directly:

```sh
buf lint
```

### Building

```sh
make build
```

### Publishing

```sh
make push
```
