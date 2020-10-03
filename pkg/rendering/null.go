// build js,wasm

package rendering

import (
	"github.com/GodsBoss/gggg/pkg/dom"
)

// Null implements rendering, but renders nothing.
type Null struct{}

func (n *Null) Render(output *dom.Context2D) {}
