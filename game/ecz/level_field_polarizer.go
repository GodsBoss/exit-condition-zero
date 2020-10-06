package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type polarizer struct {
	spriteMap sprite.Map

	orientation  polarizerOrientation
	configurable bool
}

func newPolarizer(spriteMap sprite.Map, options ...polarizerOption) field {
	pol := &polarizer{
		spriteMap: spriteMap,
	}
	cf := newCommonField(pol)
	for i := range options {
		options[i](pol, cf)
	}
	return cf
}

type polarizerOption func(*polarizer, *commonField)

func asPolarizerOption(cfOpt commonFieldOption) polarizerOption {
	return func(_ *polarizer, cf *commonField) {
		cfOpt(cf)
	}
}

func horizontalPolarizer() polarizerOption {
	return func(pol *polarizer, _ *commonField) {
		pol.orientation = horizontalPolarizerOrientation{}
	}
}

func verticalPolarizer() polarizerOption {
	return func(pol *polarizer, _ *commonField) {
		pol.orientation = verticalPolarizerOrientation{}
	}
}

func configurablePolarizer() polarizerOption {
	return func(pol *polarizer, _ *commonField) {
		pol.configurable = true
	}
}

func (pol *polarizer) ImmediateHit(dir direction) (bool, []direction) {
	return pol.orientation.ImmediateHit(dir)
}

func (pol *polarizer) IsConfigurable() bool {
	return pol.configurable
}

func (pol *polarizer) Configure() {
	pol.orientation = pol.orientation.turn()
}

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
