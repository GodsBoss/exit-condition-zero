package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
	"github.com/GodsBoss/exit-condition-zero/pkg/vector/int2d"
)

type direction string

const (
	dirUp    direction = "up"
	dirRight direction = "right"
	dirDown  direction = "down"
	dirLeft  direction = "left"
)

func turnDirectionClockwise(dir direction) direction {
	return (map[direction]direction{
		dirUp:    dirRight,
		dirRight: dirDown,
		dirDown:  dirLeft,
		dirLeft:  dirUp,
	})[dir]
}

func turnDirectionCounterClockwise(dir direction) direction {
	return (map[direction]direction{
		dirUp:    dirLeft,
		dirRight: dirUp,
		dirDown:  dirRight,
		dirLeft:  dirDown,
	})[dir]
}

func oppositeDirection(dir direction) direction {
	return (map[direction]direction{
		dirUp:    dirDown,
		dirRight: dirLeft,
		dirDown:  dirUp,
		dirLeft:  dirRight,
	})[dir]
}

var directionVectors = map[direction]int2d.Vector{
	dirUp:    int2d.Up(),
	dirRight: int2d.Right(),
	dirDown:  int2d.Down(),
	dirLeft:  int2d.Left(),
}

var directionSpriteIDs = map[direction]string{
	dirUp:    "output_up",
	dirRight: "output_right",
	dirDown:  "output_down",
	dirLeft:  "output_left",
}

func createRenderableForDirections(spriteMap sprite.Map, directions []direction, x, y int, scale int) game.Renderable {
	r := make(rendering.Renderables, 0)
	for i := range directions {
		if id, ok := directionSpriteIDs[directions[i]]; ok {
			r = append(r, spriteMap.Produce(id, x, y, scale, 0))
		}
	}
	return r
}
