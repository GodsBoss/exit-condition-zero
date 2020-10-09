package ecz

import (
	"github.com/GodsBoss/gggg/pkg/vector/int2d"
)

type grid struct {
	width  int
	height int
}

func (g grid) realGridPosition(v int2d.Vector) int2d.Vector {
	return int2d.FromXY(
		wrapAround(g.width, v.X()),
		wrapAround(g.height, v.Y()),
	)
}

func wrapAround(max, val int) int {
	if val > 0 {
		return val % max
	}
	return ((val % max) + val) % max
}
