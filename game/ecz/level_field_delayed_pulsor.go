package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

// delayedPulsor is a pulsor
type delayedPulsor struct {
	spriteMap sprite.Map

	// mode determines how powering the pulsor lets pulses loose.
	mode delayPulsorMode

	directions directionsMap

	// initialPowered is the value powered will be set to at Reset().
	initialPowered bool

	// powered remembers wether the pulsor had received a pulse in the last cycle.
	powered bool

	poweredBefore bool

	powerAnimation *animation
	dirAnim        *animation
}

func newDelayedPulsor(
	spriteMap sprite.Map,
	mode delayPulsorMode,
	directions map[direction]struct{},
	initialPowered bool,
	deletable bool,
	movable bool,
) field {
	return newCommonField(
		&delayedPulsor{
			spriteMap:      spriteMap,
			mode:           mode,
			directions:     directions,
			initialPowered: initialPowered,
			poweredBefore:  initialPowered,
			powered:        initialPowered,
			powerAnimation: &animation{
				fps:    8,
				frames: 4,
			},
			dirAnim: &animation{
				fps:    2,
				frames: 2,
			},
		},
		setDeletable(deletable),
		setMovable(movable),
	)
}

func (p *delayedPulsor) Reset() {
	p.poweredBefore = p.initialPowered
	p.powered = p.initialPowered
}

func (p *delayedPulsor) ExtractOutputPulses() []direction {
	dirs := p.mode.extractOutputPulses(p.powered, p.directions)
	p.poweredBefore = p.powered
	p.powered = false
	return dirs
}

func (p *delayedPulsor) ImmediateHit(direction) (bool, []direction) {
	return true, make([]direction, 0)
}

func (p *delayedPulsor) Receive([]direction) {
	p.powered = true
}

func (p *delayedPulsor) Renderable(x, y int, scale int) game.Renderable {
	return rendering.Renderables{
		p.mode.renderable(p, x, y, scale),
		createRenderableForDirections(p.spriteMap, p.directions.Directions(), x, y, scale, p.dirAnim.frame()),
	}
}

func (p *delayedPulsor) Tick(ms int) {
	p.powerAnimation.tick(ms)
	p.dirAnim.tick(ms)
}

func (p *delayedPulsor) IsConfigurable() bool {
	return false
}

func (p *delayedPulsor) Configure() {}

type delayPulsorMode interface {
	extractOutputPulses(powered bool, directions directionsMap) []direction
	renderable(p *delayedPulsor, x, y int, scale int) game.Renderable
}

type delayPulsorModeDelayed struct{}

var _ delayPulsorMode = delayPulsorModeDelayed{}

func (mode delayPulsorModeDelayed) extractOutputPulses(powered bool, directions directionsMap) []direction {
	if !powered {
		return make([]direction, 0)
	}
	return directions.Directions()
}

func (mode delayPulsorModeDelayed) renderable(p *delayedPulsor, x, y int, scale int) game.Renderable {
	if p.poweredBefore {
		return p.spriteMap.Produce("p_delayed_pulsor_powered", x, y, scale, p.powerAnimation.frame())
	}
	return p.spriteMap.Produce("p_delayed_pulsor", x, y, scale, 0)
}

type delayPulsorModeInverted struct{}

var _ delayPulsorMode = delayPulsorModeInverted{}

func (mode delayPulsorModeInverted) extractOutputPulses(powered bool, directions directionsMap) []direction {
	if powered {
		return make([]direction, 0)
	}
	return directions.Directions()
}

func (mode delayPulsorModeInverted) renderable(p *delayedPulsor, x, y int, scale int) game.Renderable {
	if p.poweredBefore {
		return p.spriteMap.Produce("p_inverted_pulsor_powered", x, y, scale, p.powerAnimation.frame())
	}
	return p.spriteMap.Produce("p_inverted_pulsor", x, y, scale, 0)
}
