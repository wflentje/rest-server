.PHONY: generate run build test fmt vet lint security clean

generate:
	mkdir -p internal/api
	oapi-codegen -config openapi/oapi-codegen.yaml openapi/openapi.yaml

run:
	go run ./cmd/server

build:
	go build -o bin/rest-server ./cmd/server

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