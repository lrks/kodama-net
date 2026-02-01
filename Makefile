.PHONY: generate build build-all test fmt-lint all

BIN := kodama-net
COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
LDFLAGS := "\
    -X 'main.commit=$(COMMIT)' \
    -X 'main.buildAt=$(BUILD_TIME)' \
"

generate:
	go generate internal/echonetlite/mra.go

build:
	go build -ldflags $(LDFLAGS) -o bin/$(BIN) ./cmd/$(BIN)

build-all:
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -ldflags $(LDFLAGS) -o bin/$(BIN)_linux_amd64   ./cmd/$(BIN)
	GOOS=linux   GOARCH=arm64 CGO_ENABLED=0 go build -ldflags $(LDFLAGS) -o bin/$(BIN)_linux_arm64   ./cmd/$(BIN)
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -ldflags $(LDFLAGS) -o bin/$(BIN)_darwin_amd64  ./cmd/$(BIN)
	GOOS=darwin  GOARCH=arm64 CGO_ENABLED=0 go build -ldflags $(LDFLAGS) -o bin/$(BIN)_darwin_arm64  ./cmd/$(BIN)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags $(LDFLAGS) -o bin/$(BIN)_windows_amd64.exe ./cmd/$(BIN)

test:
	go test ./...

fmt-lint:
	golangci-lint fmt
	golangci-lint run

all: test lint build
