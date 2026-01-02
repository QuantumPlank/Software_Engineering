package shapes

import (
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func ShapeFactory(shapeType string, params ...float64) Shape {
	switch shapeType {
	case "rectangle":
		return &Rectangle{Width: params[0], Height: params[1]}
	case "circle":
		return &Circle{Radius: params[0]}
	default:
		return nil
	}
}
