package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

var _ State = &playing{}

type playing struct {
	spriteMap SpriteMap
}

func (p *playing) Tick(ms int) {}

func (p *playing) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (p *playing) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (p *playing) Renderables() []Renderable {
	return make([]Renderable, 0)
}
