package main

import (
	"fmt"
)

func main() {
	var name string = "Betül"
	var greeting = createGreet(name)
	fmt.Printf("%s\n", greeting)
	fmt.Println(greeting)
}

func createGreet(name string) string {
	greeting := "Selam " + name + " :)"
	return greeting
}
