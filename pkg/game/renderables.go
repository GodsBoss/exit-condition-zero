package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
)

type Renderables []Renderable

func (r Renderables) Render(output *dom.Context2D) {
	for i := range r {
		r[i].Render(output)
	}
}
