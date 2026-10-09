package person

import "fmt"

type Person struct {
	name    string
	age     int
	married bool
}

func NewPerson(
	name string,
	age int,
	married bool,
) Person {
	if name == "" {
		fmt.Println("Не указано имя!")
		return Person{}
	}

	if age < 0 {
		fmt.Println("Возраст не может быть отрицательным!")
		return Person{}
	}

	return Person{
		name:    name,
		age:     age,
		married: married,
	}
}

func (p *Person) SetName(name string) {
	if name == "" {
		fmt.Println("Не указано имя!")
		return
	}

	p.name = name
}

func (p *Person) SetAge(age int) {
	if age < 0 {
		fmt.Println("Возраст не может быть отрицательным!")
		return
	}

	p.age = age
}
