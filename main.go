package main

import "fmt"

func main() {
	foods := []string{"pomato", "pizza", "pasta"}
	fmt.Printf("%v\n", foods)
	fmt.Println(len(foods))
	foods = append(foods, "tomato")
	fmt.Printf("%v\n", foods)
	fmt.Println(len(foods))
}