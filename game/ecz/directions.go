package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/vector/int2d"
)

type direction int2d.Vector

func (dir direction) Vector() int2d.Vector {
	return int2d.Vector(dir)
}

var (
	dirUp    = direction(int2d.Up())
	dirRight = direction(int2d.Right())
	dirDown  = direction(int2d.Down())
	dirLeft  = direction(int2d.Left())
)

func turnDirectionClockwise(dir direction) direction {
	return direction(int2d.RotateClockwise(dir.Vector()))
}

func turnDirectionCounterClockwise(dir direction) direction {
	return direction(int2d.RotateCounterclockwise(dir.Vector()))
}

func oppositeDirection(dir direction) direction {
	return direction(int2d.Multiply(dir.Vector(), -1))
}

func toDirectionsMap(directions ...direction) directionsMap {
	m := make(directionsMap)
	for i := range directions {
		m[directions[i]] = struct{}{}
	}
	return m
}

type directionsMap map[direction]struct{}

func (m directionsMap) Directions() []direction {
	dirs := make([]direction, 0, len(m))
	for dir := range m {
		dirs = append(dirs, dir)
	}
	return dirs
}

func (m directionsMap) Contains(dir direction) bool {
	_, ok := m[dir]
	return ok
}

func createRenderableForDirections(spriteMap sprite.Map, directions []direction, x, y int, scale int, frame int) game.Renderable {
	var directionSpriteIDs = map[direction]string{
		dirUp:    "output_up",
		dirRight: "output_right",
		dirDown:  "output_down",
		dirLeft:  "output_left",
	}
	r := make(game.Renderables, 0)
	for i := range directions {
		if id, ok := directionSpriteIDs[directions[i]]; ok {
			r = append(r, spriteMap.Produce(id, x, y, scale, frame))
		}
	}
	return r
}
