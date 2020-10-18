package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type State interface {
	Initer
	Ticker
	KeyEventReceiver
	MouseEventReceiver
	Renderer
}

type Transition struct {
	NextState string
}

type Initer interface {
	// Init is called every time a transition to that state happened.
	Init()
}

type Ticker interface {
	Tick(ms int) *Transition
}

type KeyEventReceiver interface {
	ReceiveKeyEvent(event interaction.KeyEvent) *Transition
}

type MouseEventReceiver interface {
	ReceiveMouseEvent(event interaction.MouseEvent) *Transition
}

type Renderer interface {
	Renderables(scale int) []Renderable
}

// ToState takes a renderer and uses interface detection to create a full state.
// If renderer implements Initer, Ticker, KeyEventReceiver or MouseEventReceiver,
// the corresponding methods are called when the wrapper's methods are called.
// Missing methods are filled with NOPs.
func ToState(renderer Renderer) State {
	state := &interfaceDetectionState{
		renderablesFunc:       renderer.Renderables,
		initFunc:              nopInit,
		tickFunc:              nopTick,
		receiveKeyEventFunc:   nopReceiveKeyEvent,
		receiveMouseEventFunc: nopReceiveMouseEvent,
	}
	if initer, ok := renderer.(Initer); ok {
		state.initFunc = initer.Init
	}
	if ticker, ok := renderer.(Ticker); ok {
		state.tickFunc = ticker.Tick
	}
	if eventer, ok := renderer.(KeyEventReceiver); ok {
		state.receiveKeyEventFunc = eventer.ReceiveKeyEvent
	}
	if eventer, ok := renderer.(MouseEventReceiver); ok {
		state.receiveMouseEventFunc = eventer.ReceiveMouseEvent
	}
	return state
}

type interfaceDetectionState struct {
	initFunc              func()
	tickFunc              func(ms int) *Transition
	receiveKeyEventFunc   func(event interaction.KeyEvent) *Transition
	receiveMouseEventFunc func(event interaction.MouseEvent) *Transition
	renderablesFunc       func(scale int) []Renderable
}

func (s *interfaceDetectionState) Init() {
	s.initFunc()
}

func (s *interfaceDetectionState) Tick(ms int) *Transition {
	return s.tickFunc(ms)
}

func (s *interfaceDetectionState) ReceiveKeyEvent(event interaction.KeyEvent) *Transition {
	return s.receiveKeyEventFunc(event)
}

func (s *interfaceDetectionState) ReceiveMouseEvent(event interaction.MouseEvent) *Transition {
	return s.receiveMouseEventFunc(event)
}

func (s *interfaceDetectionState) Renderables(scale int) []Renderable {
	return s.renderablesFunc(scale)
}

func nopInit() {}

func nopTick(_ int) *Transition {
	return nil
}

func nopReceiveKeyEvent(_ interaction.KeyEvent) *Transition {
	return nil
}

func nopReceiveMouseEvent(_ interaction.MouseEvent) *Transition {
	return nil
}
