package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/rect"
)

type levels struct {
	levels []*level
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

func (ls *levels) findLevelWithCoordinates(X, Y int) (*level, bool) {
	for i := range ls.levels {
		if ls.levels[i].ContainsPointer(X, Y) {
			return ls.levels[i], true
		}
	}
	return nil, false
}

type level struct {
	X        int
	Y        int
	Hover    bool
	Selected bool
}

func (lvl level) ContainsPointer(px, py int) bool {
	return rect.FromPositionAndSize(lvl.X, lvl.Y, 48, 48).Inside(px, py)
}
