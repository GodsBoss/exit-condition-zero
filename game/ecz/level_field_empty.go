package ecz

type emptyField struct{}

var _ field = &emptyField{}

func (f *emptyField) Reset() {}

func (f *emptyField) ExtractOutputPulses() []direction {
	return make([]direction, 0)
}

func (f *emptyField) ImmediateHit(dir direction) (bool, []direction) {
	return false, []direction{dir}
}

func (f *emptyField) Receive([]direction) {}

func (f *emptyField) IsDeletable() bool {
	return false
}

func (f *emptyField) IsMovable() bool {
	return false
}
