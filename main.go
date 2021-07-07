package main

import (
	"fmt"
	"goadvanced/entity"
)

const (
	Pi       = 3.14
	Timezone = "CET"
)

func main() {
	var a = 10
	b := &a
	*b = 20
	Test(b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
	person := entity.Person{
		Name: "Razvan",
	}
	person.SetName("Andrei")
	fmt.Println(person.Name)

	var vehicle entity.Vehicle
	vehicle = &entity.Car{Wheels: 4, Drive: true}
	fmt.Println(vehicle)

	var shape entity.Shape
	shape = entity.Circle{Radius: 4}
	fmt.Println(shape.Area())

	circle, ok := shape.(entity.Rectangle)

	if !ok {
		fmt.Println("It's not a rectangle")
	}

	fmt.Println(circle.Area())

	//circumference := circle.Circumference()
	//fmt.Println(circumference)

	names := make(map[string]int)

	names = map[string]int{
		"Tom": 1,
	}

	value, _ := names["Tom"]
	fmt.Println(value, len(names))

	for key, v := range names {
		fmt.Println(key, v)
	}

	var n *int
	var m = new(int)

	fmt.Println(n, m)

	switch vehicle.(type) {
	case entity.Car:
		fmt.Println("This is a car")
	default:
		fmt.Println("Other vehicle")
	}
}

func Test(a *int) {
	*a = 40
}

func GetCircumference(shape entity.Shape) {
	circle, ok := shape.(entity.Circle)
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(circle.Circumference())
}
