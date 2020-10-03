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
						v(6, 7): &blocker{
							spriteMap: spriteMap,
							movable:   true,
						},
					}
				},
			},
		},
	}
}
