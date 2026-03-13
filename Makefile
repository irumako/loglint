LOCAL_BIN := $(CURDIR)/bin

ifeq ($(OS),Windows_NT)
EXEEXT := .exe
else
EXEEXT :=
endif

GOLANGCI_LINT := $(LOCAL_BIN)/golangci-lint$(EXEEXT)
GOFUMPT := $(LOCAL_BIN)/gofumpt$(EXEEXT)
GOIMPORTS := $(LOCAL_BIN)/goimports$(EXEEXT)
GOLINES := $(LOCAL_BIN)/golines$(EXEEXT)
LOGLINT := $(LOCAL_BIN)/loglint$(EXEEXT)

.PHONY: build deps lint linter format test custom

build:
	go build -o $(LOGLINT) ./cmd/loglint

deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	GOBIN=$(LOCAL_BIN) go install mvdan.cc/gofumpt@latest
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golines@latest
	go mod tidy
	go mod verify

lint:
	$(GOLANGCI_LINT) run

format:
	$(GOFUMPT) -w .
	$(GOIMPORTS) -w .
	$(GOLINES) -w .

test:
	go test -v -cover ./...

custom:
	$(GOLANGCI_LINT) custom
