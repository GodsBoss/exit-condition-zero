package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type onOff struct {
	spriteMap     sprite.Map
	initialIsOpen bool
	isOpen        bool
	wasOpenBefore bool
	configurable  bool
}

func newOnOff(spriteMap sprite.Map, options ...onOffOption) field {
	oo := &onOff{
		spriteMap: spriteMap,
	}
	cf := newCommonField(oo)
	for i := range options {
		options[i](oo, cf)
	}
	return cf
}

type onOffOption func(*onOff, *commonField)

func asOnOffOption(cfOpt commonFieldOption) onOffOption {
	return func(_ *onOff, cf *commonField) {
		cfOpt(cf)
	}
}

func onOffStartOpen() onOffOption {
	return func(oo *onOff, cf *commonField) {
		oo.initialIsOpen = true
		oo.isOpen = true
		oo.wasOpenBefore = true
	}
}

func configurableOnOff() onOffOption {
	return func(oo *onOff, cf *commonField) {
		oo.configurable = true
	}
}

func (oo *onOff) Reset() {
	oo.isOpen = oo.initialIsOpen
	oo.wasOpenBefore = oo.isOpen
}

func (oo *onOff) ExtractOutputPulses() []direction {
	oo.wasOpenBefore = oo.isOpen
	oo.isOpen = false
	return make([]direction, 0)
}

func (oo *onOff) ImmediateHit(dir direction) (bool, []direction) {
	if oo.wasOpenBefore {
		return false, []direction{dir}
	}
	return true, make([]direction, 0)
}

func (oo *onOff) Receive([]direction) {
	oo.isOpen = true
}

func (oo *onOff) Renderable(x, y int, scale int) game.Renderable {
	spriteID := (map[bool]string{
		true:  "on_off_open",
		false: "on_off_closed",
	})[oo.wasOpenBefore]
	return oo.spriteMap.Produce(spriteID, x, y, scale, 0)
}

func (oo *onOff) IsConfigurable() bool {
	return oo.configurable
}

func (oo *onOff) Configure() {
	oo.initialIsOpen = !oo.initialIsOpen
	oo.isOpen = oo.initialIsOpen
}

func (oo *onOff) Tick(ms int) {}
