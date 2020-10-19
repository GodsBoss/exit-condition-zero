package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/vector/int2d"
)

type board struct {
	fields map[int2d.Vector]field
}

func (b *board) Tick(ms int) {
	for v := range b.fields {
		b.fields[v].Tick(ms)
	}
}

func (b *board) init(spriteMap sprite.Map, allPositions []int2d.Vector) {
	b.fields = make(map[int2d.Vector]field)
	for i := range allPositions {
		b.fields[allPositions[i]] = &emptyField{
			spriteMap: spriteMap,
		}
	}
}
