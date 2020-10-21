// +build js,wasm

package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
)

type Renderable interface {
	Render(output *dom.Context2D)
}
