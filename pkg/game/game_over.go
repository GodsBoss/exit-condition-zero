package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

var _ State = &gameOver{}

type gameOver struct {
	spriteMap SpriteMap
}

func (over *gameOver) Tick(ms int) {}

func (over *gameOver) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (over *gameOver) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (over *gameOver) Renderables() []Renderable {
	return make([]Renderable, 0)
}
