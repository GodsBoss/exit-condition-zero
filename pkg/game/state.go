package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type state interface {
	Tick(ms int)
	ReceiveKeyEvent(event interaction.KeyEvent)
	ReceiveMouseEvent(event interaction.MouseEvent)
}
