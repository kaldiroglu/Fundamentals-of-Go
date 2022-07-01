package main

import (
	"fmt"
)

func main() {
	i := 5
	pointer := &i
	fmt.Printf("%d %d \n", i, *pointer)
	fmt.Printf("%d %d \n\n", pointer, &i)

	i = 9
	var p *int = &i // Notice the type of p
	fmt.Printf("%d %d \n", i, *p)
	fmt.Printf("%d %d \n\n", p, &i)

	i = 16
	fmt.Printf("%d %d \n", i, *p)
	fmt.Printf("%d %d \n\n", p, &i)

	*p = 7
	fmt.Printf("%d %d \n", i, *p)
	fmt.Printf("%d %d \n\n", p, &i)

	fmt.Println("After creating one more pointer.")
	q := p
	fmt.Printf("%d %d %d \n", i, *p, *q)
	fmt.Printf("%d %d %d\n\n", p, q, &i)

	fmt.Println("Notice pointers are aliases!")
	*p = 89
	fmt.Printf("%d %d %d \n", i, *p, *q)
	fmt.Printf("%d %d %d\n\n", p, q, &i)

	*q = 99
	fmt.Printf("%d %d %d \n", i, *p, *q)
	fmt.Printf("%d %d %d\n\n", p, q, &i)

	p = f(i)
	fmt.Printf("%d %d \n", i, *p)
	fmt.Printf("%d %d \n\n", p, &i)

	fmt.Println("After setting to nil.")
	p = nil
	fmt.Printf("%d %d \n", i, p)
	fmt.Printf("%d %d \n", p, &i)

}

func f(i int) *int {
	fmt.Println("in f()")
	i++
	return &i
}
