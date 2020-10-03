package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

var _ state = &levelSelect{}

type levelSelect struct {
	spriteMap SpriteMap
}

func (ls *levelSelect) Tick(ms int) {}

func (ls *levelSelect) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (ls *levelSelect) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (ls *levelSelect) Renderables() []Renderable {
	return make([]Renderable, 0)
}
