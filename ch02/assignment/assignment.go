package main

import "fmt"

func main() {
	x := 1              // named variable
	b := false          // named variable
	p := &b             // named variable
	*p = true           // indirect variable
	person := Person{}  // named variable
	person.name = "bob" // struct field

	fmt.Printf("%d %t %s\n", x, *p, person.name)

	x++
	fmt.Printf("%d\n", x)
	x--
	fmt.Printf("%d\n", x)

	y := 9
	x, y = y, x
	fmt.Printf("%d %d\n", x, y)

	x += 5
	fmt.Printf("%d\n", x)
	x -= 2
	fmt.Printf("%d\n", x)
	x *= 9
	fmt.Printf("%d\n", x)
	x /= 3
	fmt.Printf("%d\n", x)
}

type Person struct {
	name string
}
