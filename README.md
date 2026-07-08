# Go REST Reference

A reference implementation of a design-first REST API in Go using the Go standard library and modern development practices.

## Features

- OpenAPI 3 design-first development
- Code generation with `oapi-codegen`
- Standard library `net/http`
- Structured logging with `log/slog`
- Request logging middleware
- Panic recovery middleware
- Graceful shutdown
- Environment-based configuration
- Unit and integration testing
- Security scanning with `gosec`
- VS Code / WSL friendly development environment

## Project Structure

```text
cmd/
    server/            Application entry point

internal/
    api/               Generated OpenAPI code
    config/            Application configuration
    handlers/          OpenAPI handler implementations
    httpserver/        HTTP server construction
    logging/           Structured logging
    middleware/        HTTP middleware

openapi/
    OpenAPI specification

scripts/
    Development helper scripts
```

## Prerequisites

- Go 1.24+
- Make
- Git

Recommended tools:

- Visual Studio Code
- WSL (Ubuntu)
- `oapi-codegen`
- `gosec`

## Running

Create a local environment file.

```bash
cp .env.example .env
```

Run the server.

```bash
make run-dev
```

The API will be available at:

```text
http://localhost:8080
```

## Generate OpenAPI Code

```bash
make generate
```

## Quality Checks

Run formatting, tests, and security checks.

```bash
make fmt
make vet
make test
make security
```

## Design Goals

This project demonstrates modern Go service development while remaining dependency-light.

Key architectural principles include:

- OpenAPI-first API design
- Standard library networking (`net/http`)
- Constructor-based dependency injection
- Configuration via environment variables
- Separation of generated and handwritten code
- Testability through clear package boundaries
- Production-ready HTTP server configuration