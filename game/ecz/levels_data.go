package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/vector/int2d"
)

func levelsData() *levels {
	return &levels{
		levels: []*level{
			{
				X: 120,
				Y: 40,
				getFields: func() map[int2d.Vector]field {
					return make(map[int2d.Vector]field)
				},
			},
		},
	}
}
