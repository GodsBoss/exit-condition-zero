// build js,wasm

package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
)

// NullRenderable implements rendering, but renders nothing.
type NullRenderable struct{}

func (n NullRenderable) Render(output *dom.Context2D) {}
