package main

import "fmt"

type Point struct {
	x, y     float64
	distance float64
}

type Circle1 struct {
	p Point
	r int
}

type Circle2 struct {
	Point
	r int
}

func main() {
	p1 := Point{x: 3, y: 4}
	c1 := Circle1{p1, 10}
	fmt.Printf("c1: %#v\n", c1)
	fmt.Printf("c1.p: %#v\n", c1.p)
	fmt.Printf("c1.p.x: %#v c1.p.y: %#v \n", c1.p.x, c1.p.y)
	//fmt.Printf("c1.p: %#v", c1.Point) // Can't say that!

	p2 := Point{x: 3, y: 4}
	c2 := Circle2{p2, 10}
	fmt.Printf("c2: %#v\n", c2)
	//fmt.Printf("c2.p: %#v", c2.p) // Can't say that!
	fmt.Printf("c2.p: %#v", c2.Point)
	fmt.Printf("c2.Point.x: %#v c2.Point.y: %#v \n", c2.Point.x, c2.Point.y)
}
