package ecz

import (
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
