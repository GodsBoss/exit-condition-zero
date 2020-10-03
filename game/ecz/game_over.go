package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type gameOver struct {
	spriteMap sprite.Map
}

func NewGameOver(spriteMap sprite.Map) game.State {
	return &gameOver{
		spriteMap: spriteMap,
	}
}

func (over *gameOver) Tick(ms int) *game.Transition {
	return nil
}

func (over *gameOver) ReceiveKeyEvent(event interaction.KeyEvent) *game.Transition {
	return nil
}

func (over *gameOver) ReceiveMouseEvent(event interaction.MouseEvent) *game.Transition {
	if event.Type == interaction.MouseUp && event.PrimaryButton() {
		return &game.Transition{
			NextState: "title",
		}
	}
	return nil
}

func (over *gameOver) Renderables(scale int) []game.Renderable {
	return make([]game.Renderable, 0)
}
