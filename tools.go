//go:build tools
// +build tools

package tools

// See also: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

import (
	_ "go.uber.org/mock/mockgen"
)
