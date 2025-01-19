.PHONY: default
default: build lint test

.PHONY: build
build:
	go build ./...

GOLANG_TOOL_PATH_TO_BIN=$(shell go env GOPATH)
GOLANGCI_LINT_CLI_VERSION?=latest
GOLANGCI_LINT_CLI_MODULE=github.com/golangci/golangci-lint/cmd/golangci-lint
GOLANGCI_LINT_CLI=$(GOLANG_TOOL_PATH_TO_BIN)/bin/golangci-lint
$(GOLANGCI_LINT_CLI):
	go install $(GOLANGCI_LINT_CLI_MODULE)@$(GOLANGCI_LINT_CLI_VERSION)

.PHONY: lint
lint: $(GOLANGCI_LINT_CLI)
	golangci-lint run


GODOC_CLI_VERSION=latest
GODOC_CLI_MODULE=golang.org/x/tools/cmd/godoc
GODOC_CLI=$(GOLANG_TOOL_PATH_TO_BIN)/bin/godoc
$(GODOC_CLI):
	go install $(GODOC_CLI_MODULE)@$(GODOC_CLI_VERSION)

.PHONY: godoc
godoc: $(GODOC_CLI)
	@echo "Open http://localhost:6060/pkg/github.com/akm/slogopts"
	godoc -http=:6060

GO_TEST_OPTIONS?=

.PHONY: test
test:
	go test $(GO_TEST_OPTIONS) ./...
GO_COVERAGE_DIR=coverage/unit
$(GO_COVERAGE_DIR):
	mkdir -p $(GO_COVERAGE_DIR)
GO_COVERAGE_HTML?=coverage.html
GO_COVERAGE_PROFILE?=coverage.txt
$(GO_COVERAGE_PROFILE):
	$(MAKE) test-with-coverage

# See https://app.codecov.io/github/akm/go-requestid/new
.PHONY: test-with-coverage
test-with-coverage: $(GO_COVERAGE_DIR)
	go test -cover ./... -args -test.gocoverdir="$(GO_COVERAGE_DIR)"

.PHONY: test-coverage
test-coverage: $(GO_COVERAGE_PROFILE)
	go tool covdata percent -i=$(GO_COVERAGE_DIR) -o $(GO_COVERAGE_PROFILE)
	go tool cover -html=$(GO_COVERAGE_PROFILE) -o $(GO_COVERAGE_HTML)
	@command -v open && open $(GO_COVERAGE_HTML) || echo "open $(GO_COVERAGE_HTML)"

.PHONY: clean
clean:
	rm -rf coverage
	rm -f $(GO_COVERAGE_HTML) $(GO_COVERAGE_PROFILE)
