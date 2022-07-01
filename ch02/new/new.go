package main

import "fmt"

func main() {
	p := new(int)
	fmt.Printf("%d %d\n", p, *p)

	q := new(int)
	fmt.Printf("%d %d\n", q, *q)

	fmt.Printf("%t\n", p == q)
	fmt.Printf("%t\n", *p == *q)

	s := new(string)
	fmt.Printf("%d \"%s\"\n", s, *s)
	*s = "I love Go"
	fmt.Printf("%d \"%s\"\n", s, *s)

	b := getABool()
	fmt.Printf("%t", *b)
}

func getABool() *bool {
	b := new(bool)
	*b = true
	return b
}
