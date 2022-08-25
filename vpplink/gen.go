//go:build tools

package vpplink

import (
	_ "github.com/edwarnicke/vpplink/cmd"
	_ "go.fd.io/govpp/binapi"
)

// Run using go generate -tags tools ./...
//go:generate go run github.com/edwarnicke/vpplink/cmd --binapi-package "go.fd.io/govpp/binapi"
