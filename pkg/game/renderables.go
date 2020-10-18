package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
)

// Renderables combines several Renderables into one.
type Renderables []Renderable

// Render renders all Renderables, in order.
func (r Renderables) Render(output *dom.Context2D) {
	for i := range r {
		r[i].Render(output)
	}
}
