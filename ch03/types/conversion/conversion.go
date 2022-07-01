package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	convertNumbers()
	//convertStringsToNumbers()
	//convertNumbersToStrings()
}

func convertNumbers() {
	// GOPL p. 55
	var apples int32 = 1
	var oranges int16 = 1
	//var compote int = apples + oranges // compile error
	var compote int = int(apples) + int(oranges)
	fmt.Println(compote)

	//b := apples == oranges  // Can't do this

	var b1 byte = 8
	var b2 byte = 15
	var b3 byte = b1 + b2
	//var b3 int = b1 + b2
	fmt.Println(b3)

	fp := 3.14
	var i int
	i = int(fp)
	fmt.Printf("fp: %f i: %d\n", fp, i)

	var pi1 float64 = math.Pi
	var pi2 = pi1
	//var pi3 float32 = pi1
	var pi3 float32 = float32(pi1)
	fmt.Printf("pi1: %g pi2: %g pi3: %g\n", pi1, pi2, pi3)
}

func convertStringsToNumbers() {
	i := 17
	stringI1 := fmt.Sprintf("%d", i)
	fmt.Println(stringI1)
	stringI1 += ":"
	fmt.Println(stringI1)

	stringI1 = strconv.Itoa(i)
	fmt.Println(stringI1)
	stringI1 += ":"
	fmt.Println(stringI1)
}

func convertNumbersToStrings() {
	i, _ := strconv.Atoi("1291")
	fmt.Println(i)
	i++
	fmt.Println(i)
	j, _ := strconv.ParseInt("1291", 10, 64)
	fmt.Println(j)
	j--
	fmt.Println(j)
}
