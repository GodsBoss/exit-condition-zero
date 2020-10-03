package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type levelSelect struct {
	spriteMap sprite.Map

	levels []level
}

func NewLevelSelect(spriteMap sprite.Map) game.State {
	return &levelSelect{
		spriteMap: spriteMap,
		levels: []level{
			{
				X: 120,
				Y: 40,
			},
		},
	}
}

func (ls *levelSelect) Init() {
	ls.clearHover()
}

func (ls *levelSelect) clearHover() {
	for i := range ls.levels {
		ls.levels[i].Hover = false
	}
}

func (ls *levelSelect) Tick(ms int) *game.Transition {
	return nil
}

func (ls *levelSelect) ReceiveKeyEvent(event interaction.KeyEvent) *game.Transition {
	return nil
}

func (ls *levelSelect) ReceiveMouseEvent(event interaction.MouseEvent) *game.Transition {
	if event.Type == interaction.MouseMove {
		ls.clearHover()
		for i := range ls.levels {
			if ls.levels[i].ContainsPointer(event.X, event.Y) {
				ls.levels[i].Hover = true
			}
		}
	}
	return nil
}

func (ls *levelSelect) Renderables(scale int) []game.Renderable {
	r := []game.Renderable{
		ls.spriteMap.Produce("bg_level_select", 0, 0, scale, 0),
	}
	for i := range ls.levels {
		id := "level_select_level"
		if ls.levels[i].Hover {
			id = "level_select_level_hover"
		}
		r = append(r, ls.spriteMap.Produce(id, ls.levels[i].X, ls.levels[i].Y, scale, 0))
	}
	return r
}

type level struct {
	X     int
	Y     int
	Hover bool
}

func (lvl level) ContainsPointer(px, py int) bool {
	left := lvl.X
	right := left + 48
	top := lvl.Y
	bottom := top + 48
	return px >= left && px <= right && py >= top && py <= bottom
}
