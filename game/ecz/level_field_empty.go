package ecz

type emptyField struct{}

func (f *emptyField) Reset() {}

func (f *emptyField) ExtractOutputPulses() []direction {
	return make([]direction, 0)
}

func (f *emptyField) IsHit(direction) (bool, []direction) {
	return false, nil
}
