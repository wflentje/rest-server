.PHONY: generate run build test fmt vet lint security clean

generate:
	mkdir -p internal/api
	oapi-codegen -config openapi/oapi-codegen.yaml openapi/openapi.yaml

run:
	go run ./cmd/server

build:
	go build -o bin/rest-server ./cmd/server

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/rest-server.exe ./cmd/server

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint: fmt vet

security:
	gosec ./...

clean:
	rm -rf bin

run-dev:
	./scripts/dev.sh

# Windows (requires PowerShell)
run-windows:
	powershell -ExecutionPolicy Bypass -File scripts/dev.ps1
