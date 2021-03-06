package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/text"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type levelSelect struct {
	spriteMap sprite.Map

	levels *levels
}

func newLevelSelect(spriteMap sprite.Map, levels *levels) *levelSelect {
	return &levelSelect{
		spriteMap: spriteMap,
		levels:    levels,
	}
}

func (ls *levelSelect) Init() {
	ls.levels.unselectLevels()
	ls.levels.unhoverLevels()
}

func (ls *levelSelect) ReceiveMouseEvent(event interaction.MouseEvent) *game.Transition {
	if event.Type == interaction.MouseMove {
		ls.levels.unhoverLevels()
		_, lvl, ok := ls.levels.findLevelWithCoordinates(event.X, event.Y)
		if ok {
			lvl.Hover = true
		}
	}
	if event.Type == interaction.MouseDown {
		ls.levels.unselectLevels()
		_, lvl, ok := ls.levels.findLevelWithCoordinates(event.X, event.Y)
		if ok {
			lvl.Selected = true
		}
	}
	if event.Type == interaction.MouseUp {
		i, lvl, ok := ls.levels.findLevelWithCoordinates(event.X, event.Y)
		if ok && lvl.Selected {
			ls.levels.selectedLevel = i
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
	levelIndexToShowNameFor := -1
	for i := range ls.levels.levels {
		lvl := ls.levels.levels[i]
		id := "level_select_level"
		if lvl.Hover {
			id = "level_select_level_hover"
			if levelIndexToShowNameFor == -1 {
				levelIndexToShowNameFor = i
			}
		}
		if lvl.Selected {
			id = "level_select_level_selected"
			levelIndexToShowNameFor = i
		}
		r = append(r, ls.spriteMap.Produce(id, lvl.X, lvl.Y, scale, 0))
		if lvl.Tutorial {
			r = append(r, ls.spriteMap.Produce("level_select_tutorial_marker", lvl.X, lvl.Y, scale, 0))
		}
	}
	r = append(
		r,
		text.New(
			ls.spriteMap,
			"Select a level!\n't' marks tutorial\nlevels. Difficulty\nrises from left\nto right and\nfrom top to bottom.",
			201,
			5,
			scale,
		),
		text.New(
			ls.spriteMap,
			"There is no final  \nend to the game,\njust play some\nof the levels.",
			201,
			49,
			scale,
		),
	)
	if levelIndexToShowNameFor >= 0 {
		r = append(
			r,
			text.New(ls.spriteMap, ls.levels.levels[levelIndexToShowNameFor].getName(), 4, 230, scale),
		)
	}
	return r
}
