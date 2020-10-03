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

func (over *gameOver) Tick(ms int) *game.Transition {
	return nil
}

func (over *gameOver) ReceiveKeyEvent(event interaction.KeyEvent) *game.Transition {
	return nil
}

func (over *gameOver) ReceiveMouseEvent(event interaction.MouseEvent) *game.Transition {
	return nil
}

func (over *gameOver) Renderables() []game.Renderable {
	return make([]game.Renderable, 0)
}
