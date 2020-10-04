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
				Name: "Introduction",
				X:    10,
				Y:    10,
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
						v(3, 9): newPulsor(spriteMap, toDirectionsMap(dirRight)),
						v(1, 8): newBlocker(spriteMap, asBlockerOption(makeMovable())),
						v(6, 9): newFreeField(spriteMap),
						v(9, 9): newExitCondition(spriteMap),
					}
				},
			},
			{
				Name: "Swapping",
				X:    30,
				Y:    10,
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
						v(1, 9): newPulsor(spriteMap, toDirectionsMap(dirUp, dirRight)),
						v(1, 7): newPolarizer(spriteMap, verticalPolarizer(), asPolarizerOption(makeMovable())),
						v(1, 5): newExitCondition(spriteMap),
						v(3, 9): newPolarizer(spriteMap, horizontalPolarizer(), asPolarizerOption(makeMovable())),
						v(5, 9): newExitCondition(spriteMap),
					}
				},
			},
			{
				Name: "Removal",
				X:    50,
				Y:    10,
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
						v(3, 7): newPulsor(spriteMap, toDirectionsMap(dirUp)),
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
				Name: "Mirrors",
				X:    70,
				Y:    10,
				Texts: []levelText{
					{
						X: 25,
						Y: 190,
						Content: levelTextContent(
							"Mirrors reflect pulses. Use this",
							"to your advantage. One of them",
							"must be configured to do this.",
							"Use the right-most button in the",
							"top-right corner of the screen.",
						),
					},
				},
				Tutorial: true,
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(5, 5):  newPulsor(spriteMap, toDirectionsMap(dirRight, dirUp, dirLeft)),
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
				Name: "Half-mirrors",
				X:    90,
				Y:    10,
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
						v(10, 5): newPulsor(spriteMap, toDirectionsMap(dirLeft)),
						v(3, 5):  newPulsor(spriteMap, toDirectionsMap(dirRight)),
						v(5, 5):  newHalfMirror(spriteMap, halfMirrorRotation(2), configurableHalfMirror()),
						v(5, 9):  newExitCondition(spriteMap),
						v(5, 3):  newDelayedPulsor(spriteMap, toDirectionsMap(dirRight), withInvertedPulsorMode()),
						v(8, 3):  newExitCondition(spriteMap),
					}
				},
			},
			{
				Name:     "Wrap-around",
				X:        110,
				Y:        10,
				Tutorial: true,
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
						v(8, 2): newPulsor(spriteMap, toDirectionsMap(dirRight)),
						v(2, 2): newHalfMirror(spriteMap, configurableHalfMirror()),
						v(2, 7): newExitCondition(spriteMap),
					}
				},
			},
			{
				Name:     "Turn, turn, turn around!",
				X:        130,
				Y:        10,
				Tutorial: true,
				Texts: []levelText{
					{
						X: 15,
						Y: 15,
						Content: levelTextContent(
							"Rotators turn pulses",
							"around. They come in two",
							"flavors: clockwise and",
							"counter-clockwise.",
						),
					},
				},
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(3, 10): newPulsor(spriteMap, toDirectionsMap(dirUp)),
						v(3, 5):  newRotator(spriteMap, withCounterClockwiseRotation(), configurableRotator()),
						v(8, 5):  newRotator(spriteMap, configurableRotator()),
						v(1, 5):  newExitCondition(spriteMap),
						v(8, 9):  newExitCondition(spriteMap),
						v(8, 1):  newBlocker(spriteMap),
					}
				},
			},
			{
				Name:     "See you later!",
				X:        150,
				Y:        10,
				Tutorial: true,
				Texts:    []levelText{},
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(2, 8): newPulsor(spriteMap, toDirectionsMap(dirUp)),
						v(2, 5): newDelayedPulsor(spriteMap, toDirectionsMap(dirRight)),
						v(6, 5): newDelayedPulsor(spriteMap, toDirectionsMap(dirDown)),
						v(6, 8): newFreeField(spriteMap),
						v(9, 8): newDelayedPulsor(spriteMap, toDirectionsMap(dirUp), withInvertedPulsorMode()),
						v(9, 4): newExitCondition(spriteMap),
						v(4, 3): newDelayedPulsor(spriteMap, toDirectionsMap(dirRight), asDelayedPulsorOption(makeMovable())),
					}
				},
			},
			{
				Name:     "Mirror blocks",
				X:        170,
				Y:        10,
				Tutorial: true,
				Texts: []levelText{
					{
						X: 15,
						Y: 160,
						Content: levelTextContent(
							"Mirror blocks have mirrors to some",
							"or all of their sides. They reflect",
							"pulses straight back.",
						),
					},
				},
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(1, 1): newPulsor(spriteMap, toDirectionsMap(dirRight)),
						v(4, 1): newRotator(spriteMap),
						v(4, 4): newMirrorBlock(spriteMap, toDirectionsMap(dirUp, dirRight), asMirrorBlockOption(makeDeletable())),
						v(4, 6): newBlocker(spriteMap),
						v(8, 1): newExitCondition(spriteMap),
					}
				},
			},
			{
				Name:     "Open thyself!",
				X:        10,
				Y:        30,
				Tutorial: true,
				Texts: []levelText{
					{
						X: 70,
						Y: 15,
						Content: levelTextContent(
							"Open/Closed gates open when",
							"hit with pulses. They will",
							"close again on the next",
							"cycle.",
						),
					},
				},
				getFields: func(spriteMap sprite.Map) map[int2d.Vector]field {
					return map[int2d.Vector]field{
						v(1, 1): newExitCondition(spriteMap),
						v(1, 9): newExitCondition(spriteMap),
						v(1, 5): newPulsor(spriteMap, toDirectionsMap(dirUp, dirDown)),
						v(1, 3): newOnOff(spriteMap, onOffStartOpen(), asOnOffOption(makeMovable())),
						v(1, 7): newOnOff(spriteMap, asOnOffOption(makeMovable())),
						v(0, 4): newBlocker(spriteMap, asBlockerOption(makeMovable())),
					}
				},
			},
		},
	}
}

func levelTextContent(lines ...string) string {
	return strings.Join(lines, "\n")
}
