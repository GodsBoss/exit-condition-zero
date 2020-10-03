package rendering

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"

	"github.com/GodsBoss/gggg/pkg/dom"
)

type Renderables []game.Renderable

func (r Renderables) Render(output *dom.Context2D) {
	for i := range r {
		r[i].Render(output)
	}
}
