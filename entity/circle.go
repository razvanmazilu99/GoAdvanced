package entity

import "math"

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (circle Circle) Circumference() float64 {
	return math.Pi * 2 * circle.Radius
}