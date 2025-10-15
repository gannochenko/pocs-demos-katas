# Library Package

This directory contains shared libraries and generated code for the project.

## Structure

- `proto/` - Generated Protocol Buffer code
- `openapi/` - Generated OpenAPI/HTTP REST code
- `pkg/` - Shared Go packages

## Code Generation

### Protocol Buffers

Generate Protocol Buffer code:

```bash
make generate_protobuf
```

This generates Go code from `.proto` files in `contracts/protobuf/` into `lib/proto/`.

### OpenAPI

Generate OpenAPI/HTTP REST code:

```bash
make generate_openapi
```

This generates Go code from the OpenAPI specification in `contracts/openapi/task-api.yaml` into `lib/openapi/`.

The generated code includes:

- Type definitions for requests/responses
- Server interface for implementing handlers
- OpenAPI specification embedding

### Generate All

Generate both Protocol Buffer and OpenAPI code:

```bash
make generate
```

## Usage

After generation, you can implement the OpenAPI server interface in your handlers. See `pkg/task/handler.go` for an example implementation.

The generated code provides:

- Type-safe request/response structures
- Server interface that you must implement
- HTTP routing and validation
- OpenAPI specification embedding for documentation

## Packages

- `pkg/logger` - Structured logging utilities
- `pkg/validator` - Input validation helpers
- `pkg/utils` - General utility functions

## Development

```bash
go mod tidy
go test ./...
```
