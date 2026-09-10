// Package display defines the operator-facing GUI surface (noVNC, none, later CDP).
package display

import (
	"context"
	"fmt"

	"github.com/lkmavi/osg-core"
)

// Mode selects how the agent GUI is exposed to the operator.
type Mode string

const (
	ModeNone  Mode = "none"
	ModeNoVNC Mode = "novnc"
)

// Config is host-side display options.
type Config struct {
	Mode Mode
	Port int // host loopback port; 0 = allocate
}

// Stack starts and stops a display backend for one sandbox.
type Stack interface {
	Start(ctx context.Context, cfg Config) (url string, err error)
	Stop(ctx context.Context) error
}

// None is a no-op display stack.
type None struct{}

// Start returns an empty URL.
func (None) Start(_ context.Context, _ Config) (string, error) { return "", nil }

// Stop is a no-op.
func (None) Stop(_ context.Context) error { return nil }

// ErrUnsupported is returned for modes not compiled in.
var ErrUnsupported = fmt.Errorf("display: %w", core.ErrNotImplemented)
