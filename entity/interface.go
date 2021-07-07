package entity

type Vehicle interface {
	NoOfWheels() int
	CanDrive() bool
}

type Tyre interface {
	Material()
}

type Train interface {
	Vehicle
	Tyre
}

type Shape interface {
	Area() float64
}

type ExpressTrain struct {
	Wheels int
}
