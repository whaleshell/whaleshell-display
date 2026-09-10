// Package novnc starts Xvfb + WM + VNC/websockify for Ubuntu :gui images.
package novnc

import (
	"context"
	"fmt"

	"github.com/lkmavi/osg-core"
	"github.com/lkmavi/osg-display"
)

// Stack is a stub noVNC backend (P6).
type Stack struct{}

// Start is not implemented yet.
func (Stack) Start(_ context.Context, _ display.Config) (string, error) {
	return "", fmt.Errorf("novnc: %w", core.ErrNotImplemented)
}

// Stop is a no-op stub.
func (Stack) Stop(_ context.Context) error { return nil }
