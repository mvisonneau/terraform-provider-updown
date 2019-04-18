NAME          := terraform-provider-updown
FILES         := $(shell git ls-files '*.go')
.DEFAULT_GOAL := help

.PHONY: setup
setup: ## Install required libraries/tools
	go get -u -v golang.org/x/tools/cmd/goimports
	go get -u -v github.com/golang/lint/golint

.PHONY: fmt
fmt: ## Format source code
	goimports -w $(FILES)

.PHONY: lint
lint: ## Run golint and go vet against the codebase
	golint -set_exit_status . updown
	go vet ./...

.PHONY: build
build: ## Build the provider
	go build .
	strip $(NAME)

.PHONY: deps
deps: ## Fetch all dependencies
	go mod vendor

.PHONY: imports
imports: ## Fixes the syntax (linting) of the codebase
	goimports -d $(FILES)

.PHONY: clean
clean: ## Remove binary if it exists
	rm -f $(NAME)

.PHONY: all
all: lint imports build ## Test and build for all supported platforms

.PHONY: help
help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
