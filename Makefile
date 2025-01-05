GOLANGCI_LINT := $(shell which golangci-lint)

# Default target
all: tidy

# Tidy: format and vet the code
tidy:
	@go fmt $$(go list ./...)
	@go vet $$(go list ./...)

# Tests: run the tests
test:
	@go test -v ./...

# Install golangci-lint only if it's not already installed
lint-install:
	@if ! [ -x "$(GOLANGCI_LINT)" ]; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi

# Lint the code using golangci-lint
# todo reuse var if possible
lint: lint-install
	$(shell which golangci-lint) run


# Phony targets
.PHONY: all tidy test lint-install lint
