package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type levelSelect struct {
	spriteMap sprite.Map

	levels *levels
}

func NewLevelSelect(spriteMap sprite.Map) game.State {
	return &levelSelect{
		spriteMap: spriteMap,
		levels:    levelsData(),
	}
}

func (ls *levelSelect) Init() {
	ls.levels.unselectLevels()
	ls.levels.unhoverLevels()
}

func (ls *levelSelect) Tick(ms int) *game.Transition {
	return nil
}

func (ls *levelSelect) ReceiveKeyEvent(event interaction.KeyEvent) *game.Transition {
	return nil
}

func (ls *levelSelect) ReceiveMouseEvent(event interaction.MouseEvent) *game.Transition {
	if event.Type == interaction.MouseMove {
		ls.levels.unhoverLevels()
		lvl, ok := ls.levels.findLevelWithCoordinates(event.X, event.Y)
		if ok {
			lvl.Hover = true
		}
	}
	if event.Type == interaction.MouseDown {
		ls.levels.unselectLevels()
		lvl, ok := ls.levels.findLevelWithCoordinates(event.X, event.Y)
		if ok {
			lvl.Selected = true
		}
	}
	if event.Type == interaction.MouseUp {
		lvl, ok := ls.levels.findLevelWithCoordinates(event.X, event.Y)
		if ok && lvl.Selected {
			return &game.Transition{
				NextState: "playing",
			}
		}
		ls.levels.unselectLevels()
	}
	return nil
}

func (ls *levelSelect) Renderables(scale int) []game.Renderable {
	r := []game.Renderable{
		ls.spriteMap.Produce("bg_level_select", 0, 0, scale, 0),
	}
	for i := range ls.levels.levels {
		lvl := ls.levels.levels[i]
		id := "level_select_level"
		if lvl.Hover {
			id = "level_select_level_hover"
		}
		if lvl.Selected {
			id = "level_select_level_selected"
		}
		r = append(r, ls.spriteMap.Produce(id, lvl.X, lvl.Y, scale, 0))
	}
	return r
}
