package ecz

import (
	"math"
)

type animation struct {
	progress float64

	fps    int
	frames int
}

func (anim *animation) tick(ms int) {
	anim.progress += float64(anim.fps) * float64(ms) / 1000.0
	if anim.progress >= float64(anim.frames) {
		anim.progress -= float64(anim.frames)
	}
}

func (anim *animation) frame() int {
	return int(math.Floor(anim.progress))
}
