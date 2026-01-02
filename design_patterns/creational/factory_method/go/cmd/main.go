package main

import (
	"fmt"

	shapes "factory_method.example/internal"
)

func main() {
	c := shapes.ShapeFactory("circle", 5.0)
	r := shapes.ShapeFactory("rectangle", 2.0, 3.0)

	fmt.Printf("Circle -> Area: %.2f, Perimeter: %.2f\n", c.Area(), c.Perimeter())
	fmt.Printf("Rectangle -> Area: %.2f, Perimeter: %.2f\n", r.Area(), r.Perimeter())

}
