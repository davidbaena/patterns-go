package main

import "math"

type Circle struct {
	radius float32
}

func NewCircle(c Circle) Circle {
	return Circle{
		radius: c.radius,
	}
}

func (c Circle) clone() Shape {
	return NewCircle(c)
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}
