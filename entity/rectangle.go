package entity

type Rectangle struct {
	Width, Length float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Length
}
