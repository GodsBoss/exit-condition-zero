package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type State interface {
	Tick(ms int) *Transition
	ReceiveKeyEvent(event interaction.KeyEvent) *Transition
	ReceiveMouseEvent(event interaction.MouseEvent) *Transition
	Renderables() []Renderable
}

type Transition struct {
	NextState string
}
