SHELL := /bin/bash

.DEFAULT_GOAL := all
.PHONY: all
all: ## Build, lint, and test the project.
all: mod inst gen build spell lint test

.PHONY: ci
ci: ## Run the full CI pipeline, including build and diff checks.
ci: all diff

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Remove build artifacts and clean project directory.
	$(call print-target)
	rm -rf dist
	rm -f coverage.*
	rm -f '"$(shell go env GOCACHE)/../golangci-lint"'
	go clean -i -cache -testcache -modcache -fuzzcache -x

.PHONY: mod
mod: ## Update and tidy Go module dependencies.
	$(call print-target)
	go mod tidy
	cd tools && go mod tidy

.PHONY: inst
inst: ## Install required Go tools from the tools directory.
	$(call print-target)
	cd tools && go install $(shell cd tools && go list -e -f '{{ join .Imports " " }}' -tags=tools)

.PHONY: gen
gen: ## Generate Go source code.
	$(call print-target)
	go generate ./...

.PHONY: build
build: ## Build the project using goreleaser for the current platform.
	$(call print-target)
	goreleaser build --clean --single-target --snapshot

.PHONY: spell
spell: ## Check and fix spelling errors in Markdown files.
	$(call print-target)
	misspell -error -locale=US -w **.md

.PHONY: lint
lint: ## Lint the project's Go code and automatically fix issues if possible.
	$(call print-target)
	golangci-lint run --fix

.PHONY: test
test: ## Run Go tests with race detection and coverage reporting.
	$(call print-target)
	go test -race -covermode=atomic -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: diff
diff: ## Check for uncommitted Git changes and fail if any are found.
	$(call print-target)
	git diff --exit-code
	RES=$$(git status --porcelain) ; if [ -n "$$RES" ]; then echo $$RES && exit 1 ; fi

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
