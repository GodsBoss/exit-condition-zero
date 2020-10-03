package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type gameOver struct{}

func (over *gameOver) Tick(ms int) {}

func (over *gameOver) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (over *gameOver) ReceiveMouseEvent(event interaction.MouseEvent) {}
