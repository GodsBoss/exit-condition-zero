package ecz

type level struct {
	X        int
	Y        int
	Hover    bool
	Selected bool
}

func (lvl level) ContainsPointer(px, py int) bool {
	left := lvl.X
	right := left + 48
	top := lvl.Y
	bottom := top + 48
	return px >= left && px <= right && py >= top && py <= bottom
}
