package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type pulsor struct {
	spriteMap sprite.Map

	directions map[direction]bool
	deletable  bool
	movable    bool
}

var _ field = &pulsor{}

func (p *pulsor) Reset() {}

func (p *pulsor) getOutputDirections() []direction {
	dirs := make([]direction, 0)
	for dir := range p.directions {
		if p.directions[dir] {
			dirs = append(dirs, dir)
		}
	}
	return dirs
}

func (p *pulsor) ExtractOutputPulses() []direction {
	return p.getOutputDirections()
}

func (p *pulsor) ImmediateHit(direction) (bool, []direction) {
	return true, make([]direction, 0)
}

func (p *pulsor) Receive([]direction) {}

func (p *pulsor) IsDeletable() bool {
	return p.deletable
}

func (p *pulsor) IsMovable() bool {
	return p.movable
}

func (p *pulsor) IsConfigurable() bool {
	return false
}

func (p *pulsor) Configure() {}

func (p *pulsor) Renderable(x, y int, scale int) game.Renderable {
	return rendering.Renderables{
		p.spriteMap.Produce("p_source", x, y, scale, 0),
		createRenderableForDirections(p.spriteMap, p.getOutputDirections(), x, y, scale),
	}
}
