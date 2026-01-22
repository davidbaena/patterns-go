package main

type Rectangle struct {
	width  int
	height int
}

func NewRectangle(r Rectangle) Rectangle {
	return Rectangle{
		width:  r.width,
		height: r.height,
	}
}

func (r Rectangle) clone() Shape {
	return NewRectangle(r)
}

func (r Rectangle) area() float32 {
	return float32(r.width * r.height)
}
