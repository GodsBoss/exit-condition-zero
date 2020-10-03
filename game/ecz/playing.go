package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type playing struct {
	spriteMap sprite.Map
}

func NewPlaying() game.State {
	return &playing{}
}

func (p *playing) Tick(ms int) {}

func (p *playing) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (p *playing) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (p *playing) Renderables() []game.Renderable {
	return make([]game.Renderable, 0)
}
