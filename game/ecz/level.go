package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/rect"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
	"github.com/GodsBoss/exit-condition-zero/pkg/vector/int2d"
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

	getFields func(sprite.Map) map[int2d.Vector]field
}

func (lvl level) ContainsPointer(px, py int) bool {
	return rect.FromPositionAndSize(lvl.X, lvl.Y, 48, 48).Inside(px, py)
}
