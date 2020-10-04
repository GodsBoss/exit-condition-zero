package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type exitConditionField struct {
	spriteMap sprite.Map

	hasBeenHit bool
	movable    bool

	anim *animation
}

func newExitCondition(spriteMap sprite.Map, options ...exitConditionOption) field {
	f := &exitConditionField{
		spriteMap: spriteMap,
		anim: &animation{
			fps:    10,
			frames: 4,
		},
	}
	for i := range options {
		options[i](f)
	}
	return f
}

type exitConditionOption func(*exitConditionField)

func moveableExitCondition() exitConditionOption {
	return func(f *exitConditionField) {
		f.movable = true
	}
}

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

func (f *exitConditionField) Configure() {}

func (f *exitConditionField) AllowsVictory() bool {
	return !f.hasBeenHit
}

func (f *exitConditionField) Renderable(x, y int, scale int) game.Renderable {
	return f.spriteMap.Produce("p_exit_condition", x, y, scale, f.anim.frame())
}

func (f *exitConditionField) Tick(ms int) {
	f.anim.tick(ms)
}
