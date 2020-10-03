package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type title struct {
	spriteMap sprite.Map
}

func NewTitle(spriteMap sprite.Map) game.State {
	return &title{
		spriteMap: spriteMap,
	}
}

func (t *title) Tick(ms int) *game.Transition {
	return nil
}

func (t *title) ReceiveKeyEvent(event interaction.KeyEvent) *game.Transition {
	return nil
}

func (t *title) ReceiveMouseEvent(event interaction.MouseEvent) *game.Transition {
	return nil
}

func (t *title) Renderables(scale int) []game.Renderable {
	return make([]game.Renderable, 0)
}
