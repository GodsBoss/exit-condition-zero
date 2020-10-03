package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type State interface {
	Tick(ms int)
	ReceiveKeyEvent(event interaction.KeyEvent)
	ReceiveMouseEvent(event interaction.MouseEvent)
	Renderables() []Renderable
}
