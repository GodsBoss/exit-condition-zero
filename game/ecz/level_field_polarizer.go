package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type polarizer struct {
	spriteMap sprite.Map

	orientation polarizerOrientation
	deletable   bool
	movable     bool
}

var _ field = &polarizer{}

func (pol *polarizer) Reset() {}

func (pol *polarizer) ExtractOutputPulses() []direction {
	return make([]direction, 0)
}

func (pol *polarizer) ImmediateHit(dir direction) (bool, []direction) {
	return pol.orientation.ImmediateHit(dir)
}

func (pol *polarizer) Receive([]direction) {}

func (pol *polarizer) IsDeletable() bool {
	return pol.deletable
}

func (pol *polarizer) IsMovable() bool {
	return pol.movable
}

func (pol *polarizer) IsConfigurable() bool {
	return false
}

func (pol *polarizer) Configure() {}

func (pol *polarizer) Renderable(x, y int, scale int) game.Renderable {
	return pol.orientation.renderable(pol.spriteMap, x, y, scale)
}

type polarizerOrientation interface {
	renderable(spriteMap sprite.Map, x, y int, scale int) game.Renderable
	ImmediateHit(direction) (bool, []direction)
	turn() polarizerOrientation
}

type horizontalPolarizerOrientation struct{}

var _ polarizerOrientation = horizontalPolarizerOrientation{}

func (orient horizontalPolarizerOrientation) ImmediateHit(dir direction) (bool, []direction) {
	if dir == dirLeft || dir == dirRight {
		return false, []direction{dir}
	}
	return true, make([]direction, 0)
}

func (orient horizontalPolarizerOrientation) renderable(spriteMap sprite.Map, x, y int, scale int) game.Renderable {
	return spriteMap.Produce("p_polarizer_horizontal", x, y, scale, 0)
}

func (orient horizontalPolarizerOrientation) turn() polarizerOrientation {
	return verticalPolarizerOrientation{}
}

type verticalPolarizerOrientation struct{}

var _ polarizerOrientation = verticalPolarizerOrientation{}

func (orient verticalPolarizerOrientation) ImmediateHit(dir direction) (bool, []direction) {
	if dir == dirUp || dir == dirDown {
		return false, []direction{dir}
	}
	return true, make([]direction, 0)
}

func (orient verticalPolarizerOrientation) renderable(spriteMap sprite.Map, x, y int, scale int) game.Renderable {
	return spriteMap.Produce("p_polarizer_vertical", x, y, scale, 0)
}

func (orient verticalPolarizerOrientation) turn() polarizerOrientation {
	return horizontalPolarizerOrientation{}
}
