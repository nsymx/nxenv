//go:build linux && amd64
// +build linux,amd64

package nxenv

import (
	"testing"
)

func TestDetectSystem(t *testing.T) {
	DetectSystem()
}

func TestClearScreen(t *testing.T) {
	ClearScreen()
}
