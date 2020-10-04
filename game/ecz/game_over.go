package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/text"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

// gameOver is the state showing that the level was finished successfully.
// A "Game over" in a failure sense does not exist.
type gameOver struct {
	spriteMap sprite.Map
}

func newGameOver(spriteMap sprite.Map) game.State {
	return &gameOver{
		spriteMap: spriteMap,
	}
}

func (over *gameOver) Init() {}

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
	return []game.Renderable{
		over.spriteMap.Produce("bg_game_over", 0, 0, scale, 0),
		text.New(over.spriteMap, "Congratulations!\nYou managed to break\nthe loop successfully!", 95, 100, scale),
	}
}
