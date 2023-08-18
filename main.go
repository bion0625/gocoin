package main

import "fmt"

type person struct {
	name string
	age  int
}

func (instance person) SayHello(){
	fmt.Printf("Hello my name is %s and I'm %d\n", instance.name, instance.age)
}

func (p person) SayKoranAge(){
	fmt.Printf("my korean age is %d\n", p.age+2)
}

func main() {
	uj := person{
		name: "uj",
		age:  32,
	}
	uj.SayHello()
	uj.SayKoranAge()
}
