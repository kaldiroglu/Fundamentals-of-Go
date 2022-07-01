package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	greeting := createGreet(name)
	fmt.Printf("%s\n", greeting)
	//fmt.Println(os.Args[0]) // Zeroth argument is the name of the running executable
}

func greet(name string) {
	fmt.Printf("Selam %s :)\n", name)
}

func createGreet(name string) string {
	return "Selam " + name + " :)"
}
