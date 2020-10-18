// build js,wasm

package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
)

// NullRenderable implements rendering, but renders nothing.
type NullRenderable struct{}

// Render does nothing.
func (n NullRenderable) Render(_ *dom.Context2D) {}
