package main

import "fmt"

func main() {
	circle1 := Circle{
		radius: 10.0,
	}

	rectangle1 := Rectangle{
		width:  10,
		height: 12,
	}

	cloneRectangle := rectangle1.clone()
	cloneCircle := circle1.clone()

	var listOfShapes = make([]Shape, 0)
	listOfShapes = append(listOfShapes, cloneCircle)
	listOfShapes = append(listOfShapes, cloneRectangle)

	for _, shape := range listOfShapes {
		fmt.Printf("area %f\n", shape.area())
	}
}
