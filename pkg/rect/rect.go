package rect

type Rectangle struct {
	left   int
	top    int
	right  int
	bottom int
}

func FromPositionAndSize(x, y int, width, height int) Rectangle {
	return FromSides(x, y, x+width, y+height)
}

func FromSides(horizontal1, vertical1, horizontal2, vertical2 int) Rectangle {
	r := Rectangle{}
	if horizontal1 < horizontal2 {
		r.left = horizontal1
		r.right = horizontal2
	} else {
		r.left = horizontal2
		r.right = horizontal1
	}
	if vertical1 < vertical2 {
		r.top = vertical1
		r.bottom = vertical2
	} else {
		r.top = vertical2
		r.bottom = vertical1
	}
	return r
}

func (r Rectangle) Inside(x, y int) bool {
	return r.left <= x && r.right >= x && r.top <= y && r.bottom >= y
}
