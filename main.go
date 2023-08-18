package main

import "fmt"

func test(name string){
	for _, letters := range name{
		fmt.Println(string(letters))
	}
}

func main() {
	name := "is my name"
	test(name)
}

