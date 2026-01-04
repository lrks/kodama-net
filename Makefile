.PHONY: build build-all test fmt-lint all

BIN := kodama-net

build:
	go build -o bin/$(BIN) ./cmd/$(BIN)

build-all:
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(BIN)_linux_amd64   ./cmd/$(BIN)
	GOOS=linux   GOARCH=arm64 CGO_ENABLED=0 go build -o bin/$(BIN)_linux_arm64   ./cmd/$(BIN)
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(BIN)_darwin_amd64  ./cmd/$(BIN)
	GOOS=darwin  GOARCH=arm64 CGO_ENABLED=0 go build -o bin/$(BIN)_darwin_arm64  ./cmd/$(BIN)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(BIN)_windows_amd64.exe ./cmd/$(BIN)

test:
	go test ./...

fmt-lint:
	golangci-lint fmt
	golangci-lint run

all: test lint build
