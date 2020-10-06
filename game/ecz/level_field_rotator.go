package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type rotator struct {
	spriteMap        sprite.Map
	counterClockwise bool
	configurable     bool
	anim             *animation
}

func newRotator(spriteMap sprite.Map, options ...rotatorOption) field {
	r := &rotator{
		spriteMap: spriteMap,
		anim: &animation{
			fps:    16,
			frames: 16,
		},
	}
	cf := newCommonField(r)
	for i := range options {
		options[i](r, cf)
	}
	return cf
}

type rotatorOption func(*rotator, *commonField)

func withCounterClockwiseRotation() rotatorOption {
	return func(r *rotator, _ *commonField) {
		r.counterClockwise = true
	}
}

func asRotatorOption(cfOpt commonFieldOption) rotatorOption {
	return func(_ *rotator, cf *commonField) {
		cfOpt(cf)
	}
}

func configurableRotator() rotatorOption {
	return func(r *rotator, _ *commonField) {
		r.configurable = true
	}
}

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

	return r.spriteMap.Produce(spriteID, x, y, scale, r.anim.frame())
}

func (r *rotator) Tick(ms int) {
	r.anim.tick(ms)
}

func (r *rotator) IsConfigurable() bool {
	return r.configurable
}

func (r *rotator) Configure() {
	r.counterClockwise = !r.counterClockwise
}
