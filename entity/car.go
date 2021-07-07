package entity

type Car struct {
	Wheels int
	Drive  bool
}

func (car Car) NoOfWheels() int {
	return car.Wheels
}

func (car Car) CanDrive() bool {
	return car.Drive
}
