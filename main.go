package main

import (
	"fmt"

	"github.com/gocoin/person"
)

func main() {
	uj := person.Person{}
	uj.SetDetails("uj",32)
	fmt.Println("main's uj", uj)
	fmt.Println(uj.Name())
}
