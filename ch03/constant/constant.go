package main

import (
	"fmt"
	"math"
)

const (
	e  = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
)

func main() {
	//constant1()
	//constant2()
	//untypedConstant1()
	untypedConstant2()
}

func constant1() {
	const i int = 7 // make it var and see the difference
	//i = 8
	//i++
	//i--
	fmt.Println(i)

	const s1 string = "selam" // make it var and see the difference
	//s1 = "heyyy?"
	//s1 += "naber?"
	fmt.Println(s1)

	fmt.Println(pi)
	r := 5.0 // Notice must be a fp
	fmt.Println(2 * pi * r)

	fmt.Println(math.Sqrt(3))
	fmt.Println(math.Pow(pi, 2))
}

func constant2() {
	const (
		a = 1
		b
		c = 2
		d
		e = 3
		f
	)
	fmt.Println(a, b, c, d, e, f)
}

func untypedConstant1() {
	const i int = 5 // Remove int and see what happens!
	fmt.Printf("%d is %T\n", i, i)

	var i1 int = i
	fmt.Println(i1)

	//var i2 int32 = i // Can't do this.
	//fmt.Println(i2)

	const s1 string = "selam"
	fmt.Printf("%s is %T\n", s1, s1)

	const j = 9
	fmt.Printf("%d is %T\n", j, j)

	const s2 = "heeyyooo"
	fmt.Printf("%s is %T\n", s1, s1)
}

func untypedConstant2() {
	var x float32 = math.Pi // math.pi is an untyped constant
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)

	fmt.Println(math.MinInt, math.MaxInt)

	// Ok, how about that???
	var myMinInt1 int32 = math.MinInt
	var myMaxInt1 int8 = math.MaxInt
	fmt.Println(myMinInt1, myMaxInt1)

	// Ok, how about that???
	var myMinInt2 int32 = int32(math.MinInt)
	var myMaxInt2 int8 = int8(math.MaxInt)
	fmt.Println(myMinInt2, myMaxInt2)
}
