package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type emptyField struct {
	spriteMap sprite.Map

	free bool
}

func newEmptyField(spriteMap sprite.Map) field {
	return &emptyField{
		spriteMap: spriteMap,
	}
}

func newFreeField(spriteMap sprite.Map) field {
	return &emptyField{
		spriteMap: spriteMap,
		free:      true,
	}
}

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

func (f *emptyField) IsConfigurable() bool {
	return false
}

func (f *emptyField) Configure() {}

func (f *emptyField) IsFree() bool {
	return f.free
}

func (f *emptyField) Renderable(x, y int, scale int) game.Renderable {
	spriteID := "p_field_blocked"
	if f.free {
		spriteID = "p_field_free"
	}
	return f.spriteMap.Produce(spriteID, x, y, scale, 0)
}

func (f *emptyField) Tick(ms int) {}
