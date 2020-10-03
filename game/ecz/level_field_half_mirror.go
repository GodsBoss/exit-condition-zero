package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type halfMirror struct {
	spriteMap    sprite.Map
	configurable bool

	// rotatedBy specifies how the half mirror is returned. It is between 0 and 3 (inclusive).
	rotatedBy int
}

func newHalfMirror(spriteMap sprite.Map, rotatedBy int, deletable, movable, configurable bool) field {
	return newCommonField(
		&halfMirror{
			spriteMap:    spriteMap,
			rotatedBy:    rotatedBy,
			configurable: configurable,
		},
		setDeletable(deletable),
		setMovable(movable),
	)
}

func (mirror *halfMirror) Reset() {}

func (mirror *halfMirror) ExtractOutputPulses() []direction {
	return make([]direction, 0)
}

func (mirror *halfMirror) ImmediateHit(dir direction) (bool, []direction) {
	for i := 0; i < mirror.rotatedBy; i++ {
		dir = turnDirectionCounterClockwise(dir)
	}
	if dir == dirLeft || dir == dirUp {
		return true, make([]direction, 0)
	}
	dir = (map[direction]direction{
		dirRight: dirUp,
		dirDown:  dirLeft,
	})[dir]
	for i := 0; i < mirror.rotatedBy; i++ {
		dir = turnDirectionClockwise(dir)
	}
	return false, []direction{dir}
}

func (mirror *halfMirror) Receive(_ []direction) {}

func (mirror *halfMirror) Renderable(x, y int, scale int) game.Renderable {
	spriteID := (map[int]string{
		0: "p_half_mirror_left_top",
		1: "p_half_mirror_right_top",
		2: "p_half_mirror_right_bottom",
		3: "p_half_mirror_left_bottom",
	})[mirror.rotatedBy]
	return mirror.spriteMap.Produce(spriteID, x, y, scale, 0)
}

func (mirror *halfMirror) IsConfigurable() bool {
	return mirror.configurable
}

func (mirror *halfMirror) Configure() {
	mirror.rotatedBy = (mirror.rotatedBy + 1) % 4
}
