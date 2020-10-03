package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type levelSelect struct {
	spriteMap sprite.Map

	levels []*level
}

func NewLevelSelect(spriteMap sprite.Map) game.State {
	return &levelSelect{
		spriteMap: spriteMap,
		levels: []*level{
			{
				X: 120,
				Y: 40,
			},
		},
	}
}

func (ls *levelSelect) Init() {
	ls.unselectLevels()
	ls.unhoverLevels()
}

func (ls *levelSelect) unselectLevels() {
	for i := range ls.levels {
		ls.levels[i].Selected = false
	}
}

func (ls *levelSelect) unhoverLevels() {
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
		ls.unhoverLevels()
		lvl, ok := ls.findLevelWithCoordinates(event.X, event.Y)
		if ok {
			lvl.Hover = true
		}
	}
	if event.Type == interaction.MouseDown {
		ls.unselectLevels()
		lvl, ok := ls.findLevelWithCoordinates(event.X, event.Y)
		if ok {
			lvl.Selected = true
		}
	}
	if event.Type == interaction.MouseUp {
		lvl, ok := ls.findLevelWithCoordinates(event.X, event.Y)
		if ok && lvl.Selected {
			return &game.Transition{
				NextState: "playing",
			}
		}
		ls.unselectLevels()
	}
	return nil
}

func (ls *levelSelect) findLevelWithCoordinates(X, Y int) (*level, bool) {
	for i := range ls.levels {
		if ls.levels[i].ContainsPointer(X, Y) {
			return ls.levels[i], true
		}
	}
	return nil, false
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
		if ls.levels[i].Selected {
			id = "level_select_level_selected"
		}
		r = append(r, ls.spriteMap.Produce(id, ls.levels[i].X, ls.levels[i].Y, scale, 0))
	}
	return r
}

type level struct {
	X        int
	Y        int
	Hover    bool
	Selected bool
}

func (lvl level) ContainsPointer(px, py int) bool {
	left := lvl.X
	right := left + 48
	top := lvl.Y
	bottom := top + 48
	return px >= left && px <= right && py >= top && py <= bottom
}
