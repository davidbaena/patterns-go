package main

type Shape interface {
	clone() Shape
	area() float32
}
