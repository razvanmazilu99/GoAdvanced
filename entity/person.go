package entity

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (person *Person) SetName(name string) {
	person.Name = name
}
