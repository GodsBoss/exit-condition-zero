package ecz

import (
	"strings"

	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
	"github.com/GodsBoss/exit-condition-zero/pkg/vector/int2d"
)

func levelsData() *levels {
	v := int2d.FromXY
	return &levels{
		levels: []*level{
			{
				X: 10,
				Y: 10,
				Texts: []levelText{
					{
						X: 5,
						Y: 5,
						Content: levelTextContent(
							"On every CPU cycle, a pulse is",
							"generated. To avoid an infinite loop,",
							"you have to make the exit condition",
							"zero. You can accomplish this by",
							"blocking all pulses going to the",
							"condition check.",
							"If you just run the level via the",
							"'play' button, it will never end.",
							"Block the pulse by moving the pulse",
							"blocker via the 'move' button (top",
							"right corner, middle one) between",
							"pulser and exit condition checker.",
						),
					},
				},
				Tutorial: true,
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(3, 9): &pulsor{
							spriteMap: spriteMap,
							directions: map[direction]bool{
								dirRight: true,
							},
						},
						v(1, 8): newBlocker(spriteMap, false, true),
						v(6, 9): newFreeField(spriteMap),
						v(9, 9): &exitConditionField{
							spriteMap: spriteMap,
						},
					}
				},
			},
			{
				X: 30,
				Y: 10,
				Texts: []levelText{
					{
						X: 5,
						Y: 5,
						Content: levelTextContent(
							"There are many different tools you",
							"can use. It's also possible to swap",
							"them if both could be moved.",
							"All exit conditions must be zero,",
							"else you cannot win!",
						),
					},
				},
				Tutorial: true,
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(1, 9): &pulsor{
							spriteMap: spriteMap,
							directions: map[direction]bool{
								dirUp:    true,
								dirRight: true,
							},
						},
						v(1, 7): &polarizer{
							spriteMap:   spriteMap,
							orientation: verticalPolarizerOrientation{},
							movable:     true,
						},
						v(1, 5): &exitConditionField{
							spriteMap: spriteMap,
						},
						v(3, 9): &polarizer{
							spriteMap:   spriteMap,
							orientation: horizontalPolarizerOrientation{},
							movable:     true,
						},
						v(5, 9): &exitConditionField{
							spriteMap: spriteMap,
						},
					}
				},
			},
			{
				X: 280,
				Y: 200,
				Texts: []levelText{
					{
						X:       100,
						Y:       12,
						Content: "Example text",
					},
				},
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
							true,
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
						v(3, 9): &pulsor{
							spriteMap: spriteMap,
							directions: map[direction]bool{
								dirUp:    true,
								dirRight: true,
								dirDown:  true,
								dirLeft:  true,
							},
						},
						v(3, 8):  newRotator(spriteMap, true, false, true, true),
						v(4, 9):  newFreeField(spriteMap),
						v(3, 10): newFreeField(spriteMap),
						v(2, 9):  newHalfMirror(spriteMap, 1, false, true, true),
					}
				},
			},
		},
	}
}

func levelTextContent(lines ...string) string {
	return strings.Join(lines, "\n")
}
