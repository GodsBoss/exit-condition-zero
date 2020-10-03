package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type levelSelect struct {
	spriteMap sprite.Map
}

func NewLevelSelect() game.State {
	return &levelSelect{}
}

func (ls *levelSelect) Tick(ms int) {}

func (ls *levelSelect) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (ls *levelSelect) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (ls *levelSelect) Renderables() []game.Renderable {
	return make([]game.Renderable, 0)
}
