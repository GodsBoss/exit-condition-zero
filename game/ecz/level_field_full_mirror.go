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

func newFullMirror(spriteMap sprite.Map, options ...fullMirrorOption) field {
	mirror := &fullMirror{
		spriteMap:   spriteMap,
		orientation: ascendingFullMirrorOrientation{},
	}

	cf := newCommonField(mirror)
	for i := range options {
		options[i](mirror, cf)
	}
	return cf
}

type fullMirrorOption func(*fullMirror, *commonField)

func asFullMirrorOption(cfOpt commonFieldOption) fullMirrorOption {
	return func(_ *fullMirror, cf *commonField) {
		cfOpt(cf)
	}
}

func ascendingMirror() fullMirrorOption {
	return func(mirror *fullMirror, _ *commonField) {
		mirror.orientation = ascendingFullMirrorOrientation{}
	}
}

func descendingMirror() fullMirrorOption {
	return func(mirror *fullMirror, _ *commonField) {
		mirror.orientation = descendingFullMirrorOrientation{}
	}
}

func configurableFullMirror() fullMirrorOption {
	return func(mirror *fullMirror, _ *commonField) {
		mirror.configurable = true
	}
}

func (mirror *fullMirror) ImmediateHit(dir direction) (bool, []direction) {
	return mirror.orientation.ImmediateHit(dir)
}

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
