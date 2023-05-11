//@Author: Dennis
//@Date: 2023/5/11

package shape

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}
