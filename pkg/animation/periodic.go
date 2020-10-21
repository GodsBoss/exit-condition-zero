package animation

import (
	"math"
)

// Periodic is a periodic animation. You can create a struct by yourself. The
// zero value is valid, but not very useful.
type Periodic struct {
	// FPS are the frames per second.
	FPS int

	// Frames are the number of frames of the animation. If zero, no animation
	// happens (kinda useless).
	Frames int

	// progress is the current state of the animation as a float. Always within
	// [0, Periodic.Frames).
	progress float64

	// SleepTime is the time (in ms) to wait after the animation finished before
	// advancing animation again.
	SleepTime int

	// remainingSleepTime is set to Periodic.SleepTime after the animation has
	// run, then shrinks to zero.
	remainingSleepTime int
}

func NewPeriodic(fps int, frames int, options ...PeriodicOption) *Periodic {
	p := &Periodic{
		FPS:    fps,
		Frames: frames,
	}
	for i := range options {
		options[i].configure(p)
	}
	return p
}

type PeriodicOption interface {
	configure(*Periodic)
}

type periodicOptionFunc func(*Periodic)

func (f periodicOptionFunc) configure(p *Periodic) {
	f(p)
}

func WithInitialProgress(progress float64) PeriodicOption {
	return periodicOptionFunc(
		func(p *Periodic) {
			p.progress = progress * float64(p.Frames)
		},
	)
}

func WithSleepTime(sleepTime int) PeriodicOption {
	return periodicOptionFunc(
		func(p *Periodic) {
			p.SleepTime = sleepTime
		},
	)
}

func WithInitialSleepTime(initialSleepTime int) PeriodicOption {
	return periodicOptionFunc(
		func(p *Periodic) {
			p.remainingSleepTime = initialSleepTime
		},
	)
}

// Tick advances the animation by the given milliseconds.
func (p *Periodic) Tick(ms int) {
	// We're in sleep mode.
	if p.remainingSleepTime > 0 {
		// Still more time to sleep than passes.
		if ms <= p.remainingSleepTime {
			p.remainingSleepTime -= ms
			return
		}
		// Slept enough and there is time left.
		ms -= p.remainingSleepTime
		p.remainingSleepTime = 0
	}
	// Advance progress.
	p.progress += float64(p.FPS) * float64(ms) / 1000.0
	if p.progress >= float64(p.Frames) {
		p.progress -= float64(p.Frames)
		p.remainingSleepTime = p.SleepTime
	}
}

// CurrentFrame returns the current frame of the animation. Returns an integer
// in [0, p.Frames).
func (p *Periodic) CurrentFrame() int {
	return int(math.Floor(p.progress))
}
