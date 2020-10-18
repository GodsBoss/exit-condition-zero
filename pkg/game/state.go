package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type State interface {
	Initer
	Ticker
	KeyEventReceiver
	MouseEventReceiver

	Renderables(scale int) []Renderable
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
