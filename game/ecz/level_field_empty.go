package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering"
)

type emptyField struct {
	free bool
}

var _ field = &emptyField{}
var _ fieldFree = &emptyField{}

func (f *emptyField) Reset() {}

func (f *emptyField) ExtractOutputPulses() []direction {
	return make([]direction, 0)
}

func (f *emptyField) ImmediateHit(dir direction) (bool, []direction) {
	return false, []direction{dir}
}

func (f *emptyField) Receive([]direction) {}

func (f *emptyField) IsDeletable() bool {
	return false
}

func (f *emptyField) IsMovable() bool {
	return false
}

func (f *emptyField) IsFree() bool {
	return f.free
}

func (f *emptyField) Renderable() game.Renderable {
	return rendering.Null{}
}
