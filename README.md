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
- Cross-platform development (Linux, WSL, and Windows)

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
    dev.sh             Linux / WSL development launcher
    dev.ps1            Windows PowerShell development launcher
```

## Prerequisites

- Go 1.24+
- Git

Optional development tools:

- GNU Make
- Visual Studio Code
- WSL (Ubuntu)
- `oapi-codegen`
- `gosec`

## Running

Create a local environment file.

### Linux / WSL

```bash
cp .env.example .env
make run-dev
```

### Windows PowerShell

```powershell
Copy-Item .env.example .env
.\scripts\dev.ps1
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
- Environment-based configuration
- Structured logging with `log/slog`
- Separation of generated and handwritten code
- Testability through clear package boundaries
- Production-ready HTTP server configuration