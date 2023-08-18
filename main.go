package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 50
	fmt.Println(*b)
}