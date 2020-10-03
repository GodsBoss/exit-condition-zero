package ecz

type blocker struct {
	deletable bool
	movable   bool
}

var _ field = &blocker{}

func (b *blocker) Reset() {}

func (b *blocker) ExtractOutputPulses() []direction {
	return make([]direction, 0)
}

func (b *blocker) IsHit(direction) (bool, []direction) {
	return true, make([]direction, 0)
}

func (b *blocker) IsDeletable() bool {
	return b.deletable
}

func (b *blocker) IsMovable() bool {
	return b.movable
}
