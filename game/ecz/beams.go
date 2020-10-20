package ecz

import (
	"math/rand"

	"github.com/GodsBoss/gggg/pkg/vector/int2d"
)

const beamAnimationSpeed = 0.1

type beamIndex struct {
	v         int2d.Vector
	d         direction
	firstHalf bool
}

type beam struct {
	age       float64
	animation float64
}

func newBeam() *beam {
	return &beam{
		animation: rand.Float64() * 4,
	}
}

func (b *beam) Decay() int {
	if b.age >= 0.4 {
		return 2
	}
	if b.age >= 0.2 {
		return 1
	}
	return 0
}
