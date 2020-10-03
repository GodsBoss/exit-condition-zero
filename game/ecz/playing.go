package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rect"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
	"github.com/GodsBoss/exit-condition-zero/pkg/vector/int2d"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type playing struct {
	spriteMap sprite.Map
	levels    *levels

	running bool
	fields  map[int2d.Vector]field
}

func newPlaying(spriteMap sprite.Map, levels *levels) game.State {
	return &playing{
		spriteMap: spriteMap,
		levels:    levels,
	}
}

func (p *playing) Init() {
	p.running = false
	p.fields = make(map[int2d.Vector]field)
	for x := 0; x < 11; x++ {
		for y := 0; y < 11; y++ {
			p.fields[int2d.FromXY(x, y)] = &emptyField{}
		}
	}

	lvlFields := p.levels.levels[p.levels.selectedLevel].getFields()
	for v := range lvlFields {
		p.fields[v] = lvlFields[v]
	}
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
