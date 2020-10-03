package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering"
)

type pulsor struct {
	directions map[direction]bool
	deletable  bool
	movable    bool
}

var _ field = &pulsor{}

func (p *pulsor) Reset() {}

func (p *pulsor) ExtractOutputPulses() []direction {
	dirs := make([]direction, 0)
	for dir := range p.directions {
		if p.directions[dir] {
			dirs = append(dirs, dir)
		}
	}
	return dirs
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

func (p *pulsor) Renderable(x, y int, scale int) game.Renderable {
	return rendering.Null{}
}
