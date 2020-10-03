package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type fullMirror struct {
	spriteMap    sprite.Map
	orientation  fullMirrorOrientation
	configurable bool
}

func newFullMirror(spriteMap sprite.Map, orientation fullMirrorOrientation, deletable, movable, configurable bool) field {
	return newCommonField(
		&fullMirror{
			spriteMap:    spriteMap,
			orientation:  orientation,
			configurable: configurable,
		},
		setDeletable(deletable),
		setMovable(movable),
	)
}

func (mirror *fullMirror) Reset() {}

func (mirror *fullMirror) ExtractOutputPulses() []direction {
	return make([]direction, 0)
}

func (mirror *fullMirror) ImmediateHit(dir direction) (bool, []direction) {
	return mirror.orientation.ImmediateHit(dir)
}

func (mirror *fullMirror) Receive([]direction) {}

func (mirror *fullMirror) IsConfigurable() bool {
	return mirror.configurable
}

func (mirror *fullMirror) Configure() {
	mirror.orientation = mirror.orientation.turn()
}

func (mirror *fullMirror) Renderable(x, y int, scale int) game.Renderable {
	return mirror.orientation.renderable(mirror.spriteMap, x, y, scale)
}

func (mirror *fullMirror) Tick(ms int) {}

type fullMirrorOrientation interface {
	renderable(spriteMap sprite.Map, x, y int, scale int) game.Renderable
	turn() fullMirrorOrientation
	ImmediateHit(direction) (bool, []direction)
}

type ascendingFullMirrorOrientation struct{}

var _ fullMirrorOrientation = ascendingFullMirrorOrientation{}

func (orient ascendingFullMirrorOrientation) turn() fullMirrorOrientation {
	return descendingFullMirrorOrientation{}
}

func (orient ascendingFullMirrorOrientation) ImmediateHit(dir direction) (bool, []direction) {
	return false, []direction{
		(map[direction]direction{
			dirUp:    dirRight,
			dirRight: dirUp,
			dirDown:  dirLeft,
			dirLeft:  dirDown,
		})[dir],
	}
}

func (orient ascendingFullMirrorOrientation) renderable(spriteMap sprite.Map, x, y int, scale int) game.Renderable {
	return spriteMap.Produce("p_full_mirror_asc", x, y, scale, 0)
}

type descendingFullMirrorOrientation struct{}

var _ fullMirrorOrientation = descendingFullMirrorOrientation{}

func (orient descendingFullMirrorOrientation) turn() fullMirrorOrientation {
	return ascendingFullMirrorOrientation{}
}

func (orient descendingFullMirrorOrientation) ImmediateHit(dir direction) (bool, []direction) {
	return false, []direction{
		(map[direction]direction{
			dirUp:    dirLeft,
			dirRight: dirDown,
			dirDown:  dirRight,
			dirLeft:  dirUp,
		})[dir],
	}
}

func (orient descendingFullMirrorOrientation) renderable(spriteMap sprite.Map, x, y int, scale int) game.Renderable {
	return spriteMap.Produce("p_full_mirror_desc", x, y, scale, 0)
}
