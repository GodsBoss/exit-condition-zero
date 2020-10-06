package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
)

type commonField struct {
	field simpleField

	resetFunc   func()
	pulsingFunc func() []direction
	receiveFunc func([]direction)
	tickFunc    func(int)

	deletable bool
	movable   bool
}

func newCommonField(field simpleField, options ...commonFieldOption) *commonField {
	f := &commonField{
		field: field,
	}
	for i := range options {
		options[i](f)
	}
	f.resetFunc = createResetFunc(field)
	f.pulsingFunc = createPulsingFunc(field)
	f.receiveFunc = createReceiveFunc(field)
	f.tickFunc = createTickFunc(field)
	return f
}

func (cf *commonField) Reset() {
	cf.resetFunc()
}

func (cf *commonField) ExtractOutputPulses() []direction {
	return cf.pulsingFunc()
}

func (cf *commonField) ImmediateHit(dir direction) (bool, []direction) {
	return cf.field.ImmediateHit(dir)
}

func (cf *commonField) Receive(dirs []direction) {
	cf.receiveFunc(dirs)
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

func (cf *commonField) Tick(ms int) {
	cf.tickFunc(ms)
}

type simpleField interface {
	configurableField

	ImmediateHit(direction) (bool, []direction)
	Renderable(x, y int, scale int) game.Renderable
	IsConfigurable() bool
}

type resettableField interface {
	Reset()
}

func createResetFunc(field simpleField) func() {
	if rf, ok := field.(resettableField); ok {
		return func() {
			rf.Reset()
		}
	}
	return nop
}

type pulsingField interface {
	ExtractOutputPulses() []direction
}

func createPulsingFunc(field simpleField) func() []direction {
	if pf, ok := field.(pulsingField); ok {
		return func() []direction {
			return pf.ExtractOutputPulses()
		}
	}
	return func() []direction {
		return make([]direction, 0)
	}
}

type receivingField interface {
	Receive([]direction)
}

func createReceiveFunc(field simpleField) func([]direction) {
	if rf, ok := field.(receivingField); ok {
		return func(dirs []direction) {
			rf.Receive(dirs)
		}
	}
	return func(_ []direction) {}
}

type tickableField interface {
	Tick(ms int)
}

func createTickFunc(field simpleField) func(int) {
	if tf, ok := field.(tickableField); ok {
		return func(ms int) {
			tf.Tick(ms)
		}
	}
	return func(_ int) {}
}

type configurableField interface {
	Configure()
}

type commonFieldOption func(*commonField)

func makeDeletable() func(*commonField) {
	return setDeletable(true)
}

func setDeletable(deletable bool) func(*commonField) {
	return func(field *commonField) {
		field.deletable = deletable
	}
}

func makeMovable() func(*commonField) {
	return setMovable(true)
}

func setMovable(movable bool) func(*commonField) {
	return func(field *commonField) {
		field.movable = movable
	}
}

func nop() {}
