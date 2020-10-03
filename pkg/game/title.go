package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

var _ state = &title{}

type title struct {
	spriteMap SpriteMap
}

func (t *title) Tick(ms int) {}

func (t *title) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (t *title) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (t *title) Renderables() []Renderable {
	return make([]Renderable, 0)
}
