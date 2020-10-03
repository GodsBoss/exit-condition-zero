package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
	"github.com/GodsBoss/exit-condition-zero/pkg/vector/int2d"
)

func levelsData() *levels {
	v := int2d.FromXY
	return &levels{
		levels: []*level{
			{
				X: 120,
				Y: 40,
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(5, 5): &emptyField{
							spriteMap: spriteMap,
							free:      true,
						},
						v(5, 3): &pulsor{
							spriteMap: spriteMap,
							directions: map[direction]bool{
								dirLeft: true,
								dirDown: true,
							},
						},
						v(5, 7): &exitConditionField{
							spriteMap: spriteMap,
						},
						v(6, 7): newBlocker(spriteMap, false, true),
						v(7, 7): newBlocker(spriteMap, true, false),
						v(1, 4): &polarizer{
							spriteMap:   spriteMap,
							orientation: horizontalPolarizerOrientation{},
							movable:     true,
						},
						v(1, 5): &polarizer{
							spriteMap:    spriteMap,
							orientation:  verticalPolarizerOrientation{},
							movable:      true,
							configurable: true,
						},
						v(2, 2): newFullMirror(spriteMap, ascendingFullMirrorOrientation{}, true, true, true),
						v(8, 4): newDelayedPulsor(
							spriteMap,
							delayPulsorModeDelayed{},
							map[direction]struct{}{
								dirUp: struct{}{},
							},
							false,
							false,
							false,
						),
						v(10, 4): newDelayedPulsor(
							spriteMap,
							delayPulsorModeInverted{},
							map[direction]struct{}{
								dirLeft: struct{}{},
							},
							false,
							false,
							false,
						),
						v(9, 8): newDelayedPulsor(
							spriteMap,
							delayPulsorModeInverted{},
							map[direction]struct{}{
								dirUp: struct{}{},
							},
							false,
							false,
							false,
						),
					}
				},
			},
		},
	}
}
