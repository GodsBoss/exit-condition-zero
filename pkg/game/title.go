package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type title struct{}

func (t *title) Tick(ms int) {}

func (t *title) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (t *title) ReceiveMouseEvent(event interaction.MouseEvent) {}
