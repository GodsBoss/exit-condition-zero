package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type State interface {
	// Init is called every time a transition to that state happened.
	Init()

	Tick(ms int) *Transition
	ReceiveKeyEvent(event interaction.KeyEvent) *Transition
	ReceiveMouseEvent(event interaction.MouseEvent) *Transition

	Renderables(scale int) []Renderable
}

type Transition struct {
	NextState string
}
