package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rect"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type playing struct {
	spriteMap sprite.Map

	running bool
}

func NewPlaying(spriteMap sprite.Map) game.State {
	return &playing{
		spriteMap: spriteMap,
	}
}

func (p *playing) Init() {
	p.running = false
}

func (p *playing) Tick(ms int) *game.Transition {
	return nil
}

func (p *playing) ReceiveKeyEvent(event interaction.KeyEvent) *game.Transition {
	return nil
}

func (p *playing) ReceiveMouseEvent(event interaction.MouseEvent) *game.Transition {
	if event.Type == interaction.MouseUp && event.PrimaryButton() {
		if rect.FromPositionAndSize(295, 215, 20, 20).Inside(event.X, event.Y) {
			return &game.Transition{
				NextState: "title",
			}
		}

		if rect.FromPositionAndSize(245, 215, 20, 20).Inside(event.X, event.Y) {
			p.toggleRun()
		}
	}
	return nil
}

func (p *playing) toggleRun() {
	if p.running {
		p.stopRunning()
	} else {
		p.startRunning()
	}
}

func (p *playing) startRunning() {
	p.running = true
}

func (p *playing) stopRunning() {
	p.running = false
}

func (p *playing) Renderables(scale int) []game.Renderable {
	r := []game.Renderable{
		p.spriteMap.Produce("bg_playing", 0, 0, scale, 0),
		p.spriteMap.Produce("playing_button_reset", 270, 215, scale, 0),
		p.spriteMap.Produce("playing_button_exit", 295, 215, scale, 0),
	}
	if p.running {
		r = append(r, p.spriteMap.Produce("playing_button_stop", 245, 215, scale, 0))
	} else {
		r = append(r, p.spriteMap.Produce("playing_button_run", 245, 215, scale, 0))
	}
	return r
}
