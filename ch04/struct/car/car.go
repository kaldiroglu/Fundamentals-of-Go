package main

import "fmt"

type Car struct {
	make, model string
	year        int
	speed       int
	distance    int
	convertible bool
}

func main() {
	var c1 Car
	fmt.Println(c1)

	c1.make = "Mercedes"
	c1.model = "G"
	c1.year = 2022
	fmt.Println(c1)
}
