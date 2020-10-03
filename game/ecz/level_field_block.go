package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type blocker struct {
	spriteMap sprite.Map

	deletable bool
	movable   bool
}

var _ field = &blocker{}

func (b *blocker) Reset() {}

func (b *blocker) ExtractOutputPulses() []direction {
	return make([]direction, 0)
}

func (b *blocker) ImmediateHit(direction) (bool, []direction) {
	return true, make([]direction, 0)
}

func (b *blocker) Receive([]direction) {}

func (b *blocker) IsDeletable() bool {
	return b.deletable
}

func (b *blocker) IsMovable() bool {
	return b.movable
}

func (b *blocker) IsConfigurable() bool {
	return false
}

func (b *blocker) Renderable(x, y int, scale int) game.Renderable {
	return b.spriteMap.Produce("p_block", x, y, scale, 0)
}
