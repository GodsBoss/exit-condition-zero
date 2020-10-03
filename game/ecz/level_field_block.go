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

func newBlocker(spriteMap sprite.Map, deletable, movable bool) field {
	return newCommonField(
		&blocker{
			spriteMap: spriteMap,
		},
		setDeletable(deletable),
		setMovable(movable),
	)
}

func (b *blocker) Reset() {}

func (b *blocker) ExtractOutputPulses() []direction {
	return make([]direction, 0)
}

func (b *blocker) ImmediateHit(direction) (bool, []direction) {
	return true, make([]direction, 0)
}

func (b *blocker) Receive([]direction) {}

func (b *blocker) IsConfigurable() bool {
	return false
}

func (b *blocker) Configure() {}

func (b *blocker) Renderable(x, y int, scale int) game.Renderable {
	return b.spriteMap.Produce("p_block", x, y, scale, 0)
}

func (b *blocker) Tick(ms int) {}
