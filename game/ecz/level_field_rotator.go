package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type rotator struct {
	spriteMap        sprite.Map
	counterClockwise bool
	configurable     bool
}

func newRotator(spriteMap sprite.Map, counterClockwise bool, deletable, movable bool, configurable bool) field {
	return newCommonField(
		&rotator{
			spriteMap:        spriteMap,
			counterClockwise: counterClockwise,
			configurable:     configurable,
		},
		setDeletable(deletable),
		setMovable(movable),
	)
}

func (r *rotator) Reset() {}

func (r *rotator) ExtractOutputPulses() []direction {
	return make([]direction, 0)
}

func (r *rotator) ImmediateHit(dir direction) (bool, []direction) {
	return false, []direction{
		(map[bool]func(direction) direction{
			true:  turnDirectionCounterClockwise,
			false: turnDirectionClockwise,
		})[r.counterClockwise](dir),
	}
}

func (r *rotator) Receive([]direction) {}

func (r *rotator) Renderable(x, y int, scale int) game.Renderable {
	spriteID := (map[bool]string{
		true:  "p_rotator_counter_clockwise",
		false: "p_rotator_clockwise",
	})[r.counterClockwise]

	return r.spriteMap.Produce(spriteID, x, y, scale, 0)
}

func (r *rotator) Tick(ms int) {}

func (r *rotator) IsConfigurable() bool {
	return r.configurable
}

func (r *rotator) Configure() {
	r.counterClockwise = !r.counterClockwise
}
