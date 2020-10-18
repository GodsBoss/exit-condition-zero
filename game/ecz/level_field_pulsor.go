package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type pulsor struct {
	spriteMap sprite.Map

	directions directionsMap

	anim    *animation
	dirAnim *animation
}

func newPulsor(spriteMap sprite.Map, directions directionsMap, options ...pulsorOption) field {
	p := &pulsor{
		spriteMap:  spriteMap,
		directions: directions,
		anim: &animation{
			fps:    10,
			frames: 6,
		},
		dirAnim: &animation{
			fps:    2,
			frames: 2,
		},
	}
	cf := newCommonField(p)
	for i := range options {
		options[i](p, cf)
	}
	return cf
}

type pulsorOption func(*pulsor, *commonField)

func asPulsorOption(cfOpt commonFieldOption) pulsorOption {
	return func(_ *pulsor, cf *commonField) {
		cfOpt(cf)
	}
}

func (p *pulsor) ExtractOutputPulses() []direction {
	return p.directions.Directions()
}

func (p *pulsor) ImmediateHit(direction) (bool, []direction) {
	return true, make([]direction, 0)
}

func (p *pulsor) Receive([]direction) {}

func (p *pulsor) IsConfigurable() bool {
	return false
}

func (p *pulsor) Configure() {}

func (p *pulsor) Renderable(x, y int, scale int) game.Renderable {
	return game.Renderables{
		p.spriteMap.Produce("p_pulsor", x, y, scale, p.anim.frame()),
		createRenderableForDirections(p.spriteMap, p.directions.Directions(), x, y, scale, p.dirAnim.frame()),
	}
}

func (p *pulsor) Tick(ms int) {
	p.anim.tick(ms)
	p.dirAnim.tick(ms)
}
