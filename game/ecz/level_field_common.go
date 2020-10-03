package ecz

import "github.com/GodsBoss/exit-condition-zero/pkg/game"

type commonField struct {
	field simpleField

	deletable bool
	movable   bool
}

func newCommonField(field simpleField, options ...commonFieldOption) field {
	f := &commonField{
		field: field,
	}
	for i := range options {
		options[i](f)
	}
	return f
}

var _ field = &commonField{}

func (cf *commonField) Reset() {
	cf.field.Reset()
}

func (cf *commonField) ExtractOutputPulses() []direction {
	return cf.field.ExtractOutputPulses()
}

func (cf *commonField) ImmediateHit(dir direction) (bool, []direction) {
	return cf.field.ImmediateHit(dir)
}

func (cf *commonField) Receive(dirs []direction) {
	cf.field.Receive(dirs)
}

func (cf *commonField) IsDeletable() bool {
	return cf.deletable
}

func (cf *commonField) IsMovable() bool {
	return cf.movable
}

func (cf *commonField) IsConfigurable() bool {
	return cf.field.IsConfigurable()
}

func (cf *commonField) Configure() {
	cf.field.Configure()
}

func (cf *commonField) Renderable(x, y int, scale int) game.Renderable {
	return cf.field.Renderable(x, y, scale)
}

func (cf *commonField) Tick(ms int) {}

type simpleField interface {
	Reset()
	ExtractOutputPulses() []direction
	ImmediateHit(direction) (bool, []direction)
	Receive([]direction)
	Renderable(x, y int, scale int) game.Renderable
	IsConfigurable() bool
	Configure()
}

type commonFieldOption func(*commonField)

func setDeletable(deletable bool) func(*commonField) {
	return func(field *commonField) {
		field.deletable = deletable
	}
}

func setMovable(movable bool) func(*commonField) {
	return func(field *commonField) {
		field.movable = movable
	}
}
