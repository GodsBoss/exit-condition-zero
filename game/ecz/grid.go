package ecz

import (
	"github.com/GodsBoss/gggg/pkg/vector/int2d"
)

type grid struct {
	width  int
	height int
}

func (g grid) allPositions() []int2d.Vector {
	positions := make([]int2d.Vector, g.width*g.height)
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			positions[y*g.width+x] = int2d.FromXY(x, y)
		}
	}
	return positions
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
	return ((val % max) + max) % max
}
