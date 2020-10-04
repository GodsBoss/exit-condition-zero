package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type pulsor struct {
	spriteMap sprite.Map

	directions directionsMap

	anim    *animation
	dirAnim *animation
}

func newPulsor(
	spriteMap sprite.Map,
	directions directionsMap,
	deletable bool,
	movable bool,
) field {
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
	return newCommonField(p, setMovable(movable), setDeletable(deletable))
}

func (p *pulsor) Reset() {}

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
	return rendering.Renderables{
		p.spriteMap.Produce("p_pulsor", x, y, scale, p.anim.frame()),
		createRenderableForDirections(p.spriteMap, p.directions.Directions(), x, y, scale, p.dirAnim.frame()),
	}
}

func (p *pulsor) Tick(ms int) {
	p.anim.tick(ms)
	p.dirAnim.tick(ms)
}
