NAME                := terraform-provider-updown
VERSION             := $(shell git describe --tags --abbrev=1)
FILES               := $(shell git ls-files 'updown/*.go')
DEV_REPOSITORY_PATH := local.dev/mvisonneau/updown
DEV_VERSION         := 0.0.1
OS_ARCH             := linux_amd64

.DEFAULT_GOAL := help

.PHONY: setup
setup: ## Install required libraries/tools for build tasks
	@command -v goimports 2>&1 >/dev/null    || go install golang.org/x/tools/cmd/goimports@latest
	@command -v gosec 2>&1 >/dev/null        || go install github.com/securego/gosec/v2/cmd/gosec@v2.9.6
	@command -v goveralls 2>&1 >/dev/null    || go install github.com/mattn/goveralls@latest
	@command -v ineffassign 2>&1 >/dev/null  || go install github.com/gordonklaus/ineffassign@v0.0.0-20210914165742-4cc7213b9bc8
	@command -v misspell 2>&1 >/dev/null     || go install github.com/client9/misspell/cmd/misspell@v0.3.4
	@command -v revive 2>&1 >/dev/null       || go install github.com/mgechev/revive@v1.1.3
	@command -v tfplugindocs 2>&1 >/dev/null || go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.4.0

.PHONY: fmt
fmt: setup ## Format source code
	goimports -w $(FILES)

.PHONY: lint
lint: revive vet goimports ineffassign misspell gosec ## Run all lint related tests against the codebase

.PHONY: revive
revive: setup ## Test code syntax with revive
	revive -config .revive.toml $(FILES)

.PHONY: vet
vet: ## Test code syntax with go vet
	go vet ./...

.PHONY: goimports
goimports: setup ## Test code syntax with goimports
	goimports -d $(FILES) > goimports.out
	@if [ -s goimports.out ]; then cat goimports.out; rm goimports.out; exit 1; else rm goimports.out; fi

.PHONY: ineffassign
ineffassign: setup ## Test code syntax for ineffassign
	ineffassign ./...

.PHONY: misspell
misspell: setup ## Test code with misspell
	misspell -error $(FILES)

.PHONY: gosec
gosec: setup ## Test code for security vulnerabilities
	gosec ./...

.PHONY: test
test: ## Run the tests against the codebase
	go test -v -race ./...

.PHONY: generate
docs: setup ## Test code with misspell
	tfplugindocs

.PHONY: build-local
build-local: ## Build the binaries using local GOOS
	go build .

.PHONY: install
install: build-local ## Build and install the provider locally
	mkdir -p ~/.terraform.d/plugins/$(DEV_REPOSITORY_PATH)/$(DEV_VERSION)/$(OS_ARCH)
	mv $(NAME) ~/.terraform.d/plugins/$(DEV_REPOSITORY_PATH)/$(DEV_VERSION)/$(OS_ARCH)

.PHONY: build
build: ## Build the binaries
	goreleaser release --snapshot --skip-publish --skip-sign --rm-dist

.PHONY: release
release: ## Build & release the binaries
	goreleaser release --rm-dist

.PHONY: publish-coveralls
publish-coveralls: setup ## Publish coverage results on coveralls
	goveralls -service drone.io -coverprofile=coverage.out

.PHONY: clean
clean: ## Remove binary if it exists
	rm -f $(NAME)

.PHONY: coverage
coverage: ## Generates coverage report
	rm -rf *.out
	go test -v ./... -coverpkg=./... -coverprofile=coverage.out

.PHONY: all
all: lint test build coverage ## Test, builds and ship package for all supported platforms

.PHONY: help
help: ## Displays this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
