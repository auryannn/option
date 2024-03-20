//go:build tools
// +build tools

//
//nolint:go list
package tools

// Manage tool dependencies via go.mod.
//
// https://go.dev/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
// https://github.com/golang/go/issues/25922
//nolint:compiler
import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
)
