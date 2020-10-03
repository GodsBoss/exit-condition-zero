package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type gameOver struct {
	spriteMap sprite.Map
}

func NewGameOver() game.State {
	return &gameOver{}
}

func (over *gameOver) Tick(ms int) {}

func (over *gameOver) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (over *gameOver) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (over *gameOver) Renderables() []game.Renderable {
	return make([]game.Renderable, 0)
}
