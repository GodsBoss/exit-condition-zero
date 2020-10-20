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

func (b *beam) Tick(ms int) {
	b.animation += beamAnimationSpeed
	if b.animation >= 4.0 {
		b.animation -= 4.0
	}
	b.age += float64(ms) / 1000
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

type beams struct {
	asMap   map[beamIndex]*beam
	asSlice []beamIndex
}

func (b *beams) clear() {
	b.asMap = make(map[beamIndex]*beam)
	b.asSlice = make([]beamIndex, 0)
}

func (b *beams) add(index beamIndex, addedBeam *beam) {
	b.asMap[index] = addedBeam
	b.asSlice = append(b.asSlice, index)
}

func (b *beams) Tick(ms int) {
	for bi := range b.asMap {
		b.asMap[bi].Tick(ms)
	}
}

var beamSpriteIDs = map[bool]map[direction]string{
	true: {
		dirUp:    "p_beam_up_1",
		dirRight: "p_beam_right_1",
		dirDown:  "p_beam_down_1",
		dirLeft:  "p_beam_left_1",
	},
	false: {
		dirUp:    "p_beam_up_2",
		dirRight: "p_beam_right_2",
		dirDown:  "p_beam_down_2",
		dirLeft:  "p_beam_left_2",
	},
}
