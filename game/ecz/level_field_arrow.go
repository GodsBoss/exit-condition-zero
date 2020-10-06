package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type arrowField struct {
	spriteMap          sprite.Map
	dir                direction
	possibleDirections []direction
}

func newArrowField(spriteMap sprite.Map, dir direction, options ...arrowFieldOption) field {
	af := &arrowField{
		spriteMap: spriteMap,
		dir:       dir,
	}
	cf := newCommonField(af)
	for i := range options {
		options[i](af, cf)
	}
	return cf
}

type arrowFieldOption func(*arrowField, *commonField)

func asArrowFieldOption(cfOpt commonFieldOption) arrowFieldOption {
	return func(_ *arrowField, cf *commonField) {
		cfOpt(cf)
	}
}

func withConfigurableDirections(dirs ...direction) arrowFieldOption {
	return func(af *arrowField, _ *commonField) {
		af.possibleDirections = dirs
	}
}

func (af *arrowField) ImmediateHit(direction) (bool, []direction) {
	return false, []direction{af.dir}
}

func (af *arrowField) Renderable(x, y int, scale int) game.Renderable {
	return af.spriteMap.Produce(
		(map[direction]string{
			dirUp:    "p_arrow_up",
			dirRight: "p_arrow_right",
			dirDown:  "p_arrow_down",
			dirLeft:  "p_arrow_left",
		})[af.dir],
		x,
		y,
		scale,
		0,
	)
}

func (af *arrowField) IsConfigurable() bool {
	return len(af.possibleDirections) > 0
}

func (af *arrowField) Configure() {
	for i := range af.possibleDirections {
		if af.possibleDirections[i] == af.dir {
			i = i + 1
			if i >= len(af.possibleDirections) {
				i = 0
			}
			af.dir = af.possibleDirections[i]
			return
		}
	}
}

func (af *arrowField) Tick(_ int) {}
