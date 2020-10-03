package int2d

type Vector struct {
	x int
	y int
}

func (v Vector) X() int {
	return v.x
}

func (v Vector) Y() int {
	return v.y
}

func Zero() Vector {
	return Vector{}
}

func FromXY(x int, y int) Vector {
	return Vector{
		x: x,
		y: y,
	}
}

func Up() Vector {
	return FromXY(0, -1)
}

func Right() Vector {
	return FromXY(1, 0)
}

func Down() Vector {
	return FromXY(0, 1)
}

func Left() Vector {
	return FromXY(-1, 0)
}

func Add(vectors ...Vector) Vector {
	v := Vector{}
	for i := range vectors {
		v.x += vectors[i].x
		v.y += vectors[i].y
	}
	return v
}

func Multiply(v Vector, f int) Vector {
	return Vector{
		x: v.x * f,
		y: v.y * f,
	}
}

func Negate(v Vector) Vector {
	return Vector{
		x: -v.x,
		y: -v.y,
	}
}
