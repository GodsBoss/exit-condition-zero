// build js,wasm

package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
)

type renderable interface {
	render(output *dom.Context2D)
}
