ARTIFACT_NAME = bookgo
VERSION = 0.3.0

ifeq ($(OS),Windows_NT)
  EXT = .exe
else
  EXT =
endif

.PHONY: build run clean test lint install help

help:
	@echo "Available commands:"
	@echo "  build    - Build the binary"
	@echo "  buildall - Build for the binary for all OS"
	@echo "  run      - Run the project"
	@echo "  clean    - Remove the bin directory"
	@echo "  test     - Run all tests"
	@echo "  lint     - Run golangci-lint"
	@echo "  install  - Install the binary in GOPATH/bin"

build:
	@go build -trimpath -ldflags="-X main.Version=$(VERSION)" -o bin/${ARTIFACT_NAME}${EXT} cmd/${ARTIFACT_NAME}/main.go

buildall:
	@GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-X main.Version=$(VERSION)" -o bin/${ARTIFACT_NAME}-linux-amd64 cmd/${ARTIFACT_NAME}/main.go
	@GOOS=windows GOARCH=amd64 go build -trimpath -ldflags="-X main.Version=$(VERSION)" -o bin/${ARTIFACT_NAME}-windows-amd64.exe cmd/${ARTIFACT_NAME}/main.go
	@GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags="-X main.Version=$(VERSION)" -o bin/${ARTIFACT_NAME}-darwin-amd64 cmd/${ARTIFACT_NAME}/main.go

run:
	@go run cmd/${ARTIFACT_NAME}/main.go

clean:
	@rm -rf bin

test:
	@go test ./...

lint:
	@golangci-lint run ./...

install:
	@go install -ldflags="-X main.Version=$(VERSION)" cmd/${ARTIFACT_NAME}/main.go