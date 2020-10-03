package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type exitConditionField struct {
	spriteMap sprite.Map

	hasBeenHit bool
	movable    bool
}

var _ field = &exitConditionField{}
var _ fieldWithVictoryCondition = &exitConditionField{}

func (f *exitConditionField) Reset() {}

func (f *exitConditionField) ExtractOutputPulses() []direction {
	f.hasBeenHit = false
	return make([]direction, 0)
}

func (f *exitConditionField) ImmediateHit(direction) (bool, []direction) {
	return true, make([]direction, 0)
}

func (f *exitConditionField) Receive([]direction) {
	f.hasBeenHit = true
}

func (f *exitConditionField) IsDeletable() bool {
	return false
}

func (f *exitConditionField) IsMovable() bool {
	return f.movable
}

func (f *exitConditionField) IsConfigurable() bool {
	return false
}

func (f *exitConditionField) AllowsVictory() bool {
	return !f.hasBeenHit
}

func (f *exitConditionField) Renderable(x, y int, scale int) game.Renderable {
	return f.spriteMap.Produce("p_sink", x, y, scale, 0)
}
