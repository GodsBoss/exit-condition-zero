package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

// delayedPulsor is a pulsor
type delayedPulsor struct {
	spriteMap sprite.Map

	// mode determines how powering the pulsor lets pulses loose.
	mode delayPulsorMode

	directions map[direction]struct{}

	// initialPowered is the value powered will be set to at Reset().
	initialPowered bool

	// powered remembers wether the pulsor had received a pulse in the last cycle.
	powered bool
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
			powered:        initialPowered,
		},
		setDeletable(deletable),
		setMovable(movable),
	)
}

func (p *delayedPulsor) Reset() {
	p.powered = p.initialPowered
}

func (p *delayedPulsor) ExtractOutputPulses() []direction {
	dirs := p.mode.extractOutputPulses(p.powered, p.directions)
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
	return p.mode.renderable(p.spriteMap, x, y, scale)
}

func (p *delayedPulsor) IsConfigurable() bool {
	return false
}

func (p *delayedPulsor) Configure() {}

type delayPulsorMode interface {
	extractOutputPulses(powered bool, directions map[direction]struct{}) []direction
	renderable(spriteMap sprite.Map, x, y int, scale int) game.Renderable
}

type delayPulsorModeDelayed struct{}

var _ delayPulsorMode = delayPulsorModeDelayed{}

func (mode delayPulsorModeDelayed) extractOutputPulses(powered bool, directions map[direction]struct{}) []direction {
	if !powered {
		return make([]direction, 0)
	}
	dirs := make([]direction, 0)
	for dir := range directions {
		dirs = append(dirs, dir)
	}
	return dirs
}

func (mode delayPulsorModeDelayed) renderable(spriteMap sprite.Map, x, y int, scale int) game.Renderable {
	return spriteMap.Produce("p_delayed_pulsor", x, y, scale, 0)
}

type delayPulsorModeInverted struct{}

var _ delayPulsorMode = delayPulsorModeInverted{}

func (mode delayPulsorModeInverted) extractOutputPulses(powered bool, directions map[direction]struct{}) []direction {
	if powered {
		return make([]direction, 0)
	}
	dirs := make([]direction, 0)
	for dir := range directions {
		dirs = append(dirs, dir)
	}
	return dirs
}

func (mode delayPulsorModeInverted) renderable(spriteMap sprite.Map, x, y int, scale int) game.Renderable {
	return spriteMap.Produce("p_inverted_pulsor", x, y, scale, 0)
}
