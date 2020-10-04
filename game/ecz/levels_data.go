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
						v(3, 9): newPulsor(spriteMap, toDirectionsMap(dirRight), false, false),
						v(1, 8): newBlocker(spriteMap, asBlockerOption(makeMovable())),
						v(6, 9): newFreeField(spriteMap),
						v(9, 9): newExitCondition(spriteMap),
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
						v(1, 9): newPulsor(spriteMap, toDirectionsMap(dirUp, dirRight), false, false),
						v(1, 7): newPolarizer(spriteMap, verticalPolarizerOrientation{}, false, true, false),
						v(1, 5): newExitCondition(spriteMap),
						v(3, 9): newPolarizer(spriteMap, horizontalPolarizerOrientation{}, false, true, false),
						v(5, 9): newExitCondition(spriteMap),
					}
				},
			},
			{
				X: 50,
				Y: 10,
				Texts: []levelText{
					{
						X: 5,
						Y: 170,
						Content: levelTextContent(
							"Sometimes elements help you. Inverting",
							"pulsers send pulses when they did not",
							"receive a pulse in the last cycle. In",
							"this case, you have to remove the",
							"blocker so the inverting pulser can be",
							"disabled. Use the 'delete' button in",
							"the top right corner (left button).",
							"Also, not all elements can be moved,",
							"despite free fields being available.",
						),
					},
				},
				Tutorial: true,
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(3, 7): newPulsor(spriteMap, toDirectionsMap(dirUp), false, false),
						v(3, 3): newDelayedPulsor(
							spriteMap,
							toDirectionsMap(dirRight),
							withInvertedPulsorMode(),
						),
						v(3, 5): newBlocker(spriteMap, asBlockerOption(makeDeletable())),
						v(9, 3): newExitCondition(spriteMap),
						v(5, 5): newFreeField(spriteMap),
					}
				},
			},
			{
				X: 70,
				Y: 10,
				Texts: []levelText{
					{
						X: 25,
						Y: 210,
						Content: levelTextContent(
							"Mirrors reflect pulses. Use this",
							"to your advantage.",
						),
					},
				},
				Tutorial: true,
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(5, 5):  newPulsor(spriteMap, toDirectionsMap(dirRight, dirUp, dirLeft), false, false),
						v(7, 5):  newFreeField(spriteMap),
						v(7, 3):  newBlocker(spriteMap),
						v(9, 3):  newFullMirror(spriteMap, asFullMirrorOption(makeMovable())),
						v(10, 5): newExitCondition(spriteMap),
						v(0, 5):  newFullMirror(spriteMap, configurableFullMirror()),
						v(0, 8):  newExitCondition(spriteMap),
						v(0, 1):  newBlocker(spriteMap),
						v(5, 2):  newFullMirror(spriteMap, descendingMirror(), asFullMirrorOption(makeDeletable())),
						v(2, 2):  newExitCondition(spriteMap),
						v(5, 0):  newBlocker(spriteMap),
					}
				},
			},
			{
				X: 90,
				Y: 10,
				Texts: []levelText{
					{
						X: 15,
						Y: 15,
						Content: levelTextContent(
							"Half-mirrors reflect pulses only on",
							"two sides, the other two block them.",
						),
					},
				},
				Tutorial: true,
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(10, 5): newPulsor(spriteMap, toDirectionsMap(dirLeft), false, false),
						v(3, 5):  newPulsor(spriteMap, toDirectionsMap(dirRight), false, false),
						v(5, 5):  newHalfMirror(spriteMap, halfMirrorRotation(2), configurableHalfMirror()),
						v(5, 9):  newExitCondition(spriteMap),
						v(5, 3):  newDelayedPulsor(spriteMap, toDirectionsMap(dirRight), withInvertedPulsorMode()),
						v(8, 3):  newExitCondition(spriteMap),
					}
				},
			},
			{
				X: 70,
				Y: 30,
				Texts: []levelText{
					{
						X: 5,
						Y: 5,
						Content: levelTextContent(
							"Some elements can be configured on",
							"the spot. Use the 'configure' button",
							"in the top right corner (the right",
							"one).",
						),
					},
				},
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(5, 6):  newPulsor(spriteMap, toDirectionsMap(dirRight, dirDown, dirLeft), false, false),
						v(7, 6):  newFullMirror(spriteMap, configurableFullMirror()),
						v(7, 4):  newExitCondition(spriteMap),
						v(7, 8):  newBlocker(spriteMap),
						v(3, 6):  newHalfMirror(spriteMap, halfMirrorRotation(1), configurableHalfMirror()),
						v(3, 4):  newExitCondition(spriteMap),
						v(3, 9):  newBlocker(spriteMap),
						v(5, 7):  newRotator(spriteMap, configurableRotator()),
						v(2, 7):  newExitCondition(spriteMap),
						v(10, 7): newBlocker(spriteMap),
					}
				},
			},
			{
				X: 90,
				Y: 30,
				Texts: []levelText{
					{
						X: 100,
						Y: 100,
						Content: levelTextContent(
							"Beware of the",
							"wrap-around!",
						),
					},
				},
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(8, 2): newPulsor(spriteMap, toDirectionsMap(dirRight), false, false),
						v(2, 2): newHalfMirror(spriteMap, configurableHalfMirror()),
						v(2, 7): newExitCondition(spriteMap),
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
						v(5, 3): newPulsor(spriteMap, toDirectionsMap(dirLeft, dirDown), false, false),
						v(5, 7): newExitCondition(spriteMap),
						v(6, 7): newBlocker(spriteMap, asBlockerOption(makeMovable())),
						v(7, 7): newBlocker(spriteMap, asBlockerOption(makeDeletable())),
						v(1, 4): newPolarizer(spriteMap, horizontalPolarizerOrientation{}, false, true, false),
						v(1, 5): newPolarizer(spriteMap, verticalPolarizerOrientation{}, false, true, true),
						v(2, 2): newFullMirror(
							spriteMap,
							asFullMirrorOption(makeDeletable()),
							asFullMirrorOption(makeMovable()),
							configurableFullMirror(),
						),
						v(8, 4): newDelayedPulsor(
							spriteMap,
							toDirectionsMap(dirUp),
							delayedPulsorStartsPowered(),
						),
						v(10, 4): newDelayedPulsor(
							spriteMap,
							toDirectionsMap(dirLeft),
							withInvertedPulsorMode(),
						),
						v(9, 8): newDelayedPulsor(
							spriteMap,
							toDirectionsMap(dirUp),
							withInvertedPulsorMode(),
						),
						v(3, 9):  newPulsor(spriteMap, toDirectionsMap(dirUp, dirRight, dirDown, dirLeft), false, false),
						v(3, 8):  newRotator(spriteMap, withCounterClockwiseRotation(), asRotatorOption(makeMovable()), configurableRotator()),
						v(4, 9):  newFreeField(spriteMap),
						v(3, 10): newFreeField(spriteMap),
						v(2, 9):  newHalfMirror(spriteMap, halfMirrorRotation(1), asHalfMirrorOption(makeMovable()), configurableHalfMirror()),
					}
				},
			},
		},
	}
}

func levelTextContent(lines ...string) string {
	return strings.Join(lines, "\n")
}
