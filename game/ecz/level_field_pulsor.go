package ecz

type pulsor struct {
	directions map[direction]bool
	deletable  bool
	movable    bool
}

var _ field = &pulsor{}

func (p *pulsor) Reset() {}

func (p *pulsor) ExtractOutputPulses() []direction {
	dirs := make([]direction, 0)
	for dir := range p.directions {
		if p.directions[dir] {
			dirs = append(dirs, dir)
		}
	}
	return dirs
}

func (p *pulsor) IsHit(direction) (bool, []direction) {
	return true, make([]direction, 0)
}

func (p *pulsor) IsDeletable() bool {
	return p.deletable
}

func (p *pulsor) IsMovable() bool {
	return p.movable
}
