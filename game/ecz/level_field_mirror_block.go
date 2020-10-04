package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

type mirrorBlock struct {
	spriteMap sprite.Map
	mirrors   directionsMap
}

func newMirrorBlock(spriteMap sprite.Map, mirrors directionsMap, options ...mirrorBlockOption) field {
	bl := &mirrorBlock{
		spriteMap: spriteMap,
		mirrors:   mirrors,
	}
	cf := newCommonField(bl)
	for i := range options {
		options[i](bl, cf)
	}
	return cf
}

type mirrorBlockOption func(*mirrorBlock, *commonField)

func asMirrorBlockOption(cfOpt commonFieldOption) mirrorBlockOption {
	return func(_ *mirrorBlock, cf *commonField) {
		cfOpt(cf)
	}
}

func (bl *mirrorBlock) Reset() {}

func (bl *mirrorBlock) ExtractOutputPulses() []direction {
	return make([]direction, 0)
}

func (bl *mirrorBlock) ImmediateHit(dir direction) (bool, []direction) {
	if bl.mirrors.Contains(oppositeDirection(dir)) {
		return false, []direction{oppositeDirection(dir)}
	}
	return false, make([]direction, 0)
}

func (bl *mirrorBlock) Receive(_ []direction) {}

func (bl *mirrorBlock) Renderable(x, y int, scale int) game.Renderable {
	r := rendering.Renderables{
		bl.spriteMap.Produce("p_block_mirror_block", x, y, scale, 0),
	}
	for dir := range bl.mirrors {
		spriteID := (map[direction]string{
			dirUp:    "p_block_mirror_mirror_top",
			dirRight: "p_block_mirror_mirror_right",
			dirDown:  "p_block_mirror_mirror_bottom",
			dirLeft:  "p_block_mirror_mirror_left",
		})[dir]
		r = append(r, bl.spriteMap.Produce(spriteID, x, y, scale, 0))
	}
	return r
}

func (bl *mirrorBlock) IsConfigurable() bool {
	return false
}

func (bl *mirrorBlock) Configure() {}

func (bl *mirrorBlock) Tick(ms int) {}
