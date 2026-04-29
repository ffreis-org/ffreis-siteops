//go:build tools
// +build tools

package tools

import (
	// Blank import to ensure govulncheck is tracked in go.mod/go.sum for reproducible tool versioning.
	_ "golang.org/x/vuln/cmd/govulncheck"
)
