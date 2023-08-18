package person

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
	fmt.Println("SeeDetail's uj:", p)
}

func (p Person) Name() string {
	return p.name
}