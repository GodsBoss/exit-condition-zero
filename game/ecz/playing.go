package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type playing struct {
	spriteMap sprite.Map

	waitTime int
}

func NewPlaying(spriteMap sprite.Map) game.State {
	return &playing{
		spriteMap: spriteMap,
	}
}

func (p *playing) Init() {
	p.waitTime = 0
}

func (p *playing) Tick(ms int) *game.Transition {
	p.waitTime += ms
	if p.waitTime > 5000 {
		return &game.Transition{
			NextState: "game_over",
		}
	}
	return nil
}

func (p *playing) ReceiveKeyEvent(event interaction.KeyEvent) *game.Transition {
	return nil
}

func (p *playing) ReceiveMouseEvent(event interaction.MouseEvent) *game.Transition {
	return nil
}

func (p *playing) Renderables(scale int) []game.Renderable {
	return []game.Renderable{
		p.spriteMap.Produce("bg_playing", 0, 0, scale, 0),
	}
}
