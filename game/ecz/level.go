package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/rect"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/vector/int2d"
)

type levels struct {
	levels []*level

	// selectedLevel contains the index to an item in levels.levels.
	selectedLevel int
}

func (ls *levels) unselectLevels() {
	for i := range ls.levels {
		ls.levels[i].Selected = false
	}
}

func (ls *levels) unhoverLevels() {
	for i := range ls.levels {
		ls.levels[i].Hover = false
	}
}

func (ls *levels) findLevelWithCoordinates(X, Y int) (int, *level, bool) {
	for i := range ls.levels {
		if ls.levels[i].ContainsPointer(X, Y) {
			return i, ls.levels[i], true
		}
	}
	return -1, nil, false
}

type level struct {
	X        int
	Y        int
	Hover    bool
	Selected bool

	// Name is the level's name. It should be relatively short. The name is shown
	// in the level selection screen when hovering over the level.
	Name string

	// Tutorial determines wether this is a tutorial level. Tutorial levels are
	// easy and often have additional information in them.
	Tutorial bool

	// Texts are texts shown inside the level.
	Texts []levelText

	getFields func(sprite.Map) map[int2d.Vector]field
}

func (lvl level) getName() string {
	if lvl.Name != "" {
		return lvl.Name
	}
	return "Nameless level"
}

func (lvl level) ContainsPointer(px, py int) bool {
	return rect.FromPositionAndSize(lvl.X, lvl.Y, levelSelectLevelWidth, levelSelectLevelHeight).Inside(px, py)
}

type levelText struct {
	X int
	Y int

	Content string
}

const (
	levelSelectLevelWidth  = 20
	levelSelectLevelHeight = 20
)
