package main

import (
	"fmt"
	"math/rand"
)

var i = rand.Intn(1_000)
var j = rand.Intn(1_000)

func main() {
	arithmetic()
	//assigment()
	//unary()
	//relational()
	//bitwiseLogical()
	//shift()
}

func arithmetic() {
	i := 11
	j := 8
	k := i + j
	fmt.Println(k)

	//pi := 3.14 // float64
	//r := 5
	//circ := 2 * pi * r
	//fmt.Println(circ)
	//area := pi * r * r
	//fmt.Println(area)
}

func assigment() {
	// Assigment operators
	fmt.Println("Assigment operators")
	i = 9
	fmt.Printf("%-10s : %5d \n", "i:", i)
	i += 5
	fmt.Printf("%-10s : %5d \n", "i", i)
	i -= 12
	fmt.Printf("%-10s : %5d \n", "i", i)
	i *= 20
	fmt.Printf("%-10s : %5d \n", "i", i)
	i /= 12 // Outcome is an integer
	fmt.Printf("%-10s : %5d \n", "i", i)
	i %= 12
	fmt.Printf("%-10s : %5d \n", "i", i)
}

func unary() {
	// Unary operators
	fmt.Println("Unary operators")
	i = 9
	fmt.Printf("%-10s : %5d \n", "i:", i)
	j := +i // +i must be assigned
	fmt.Printf("%-10s : %5d \n", "j", j)
	j = -i
	fmt.Printf("%-10s : %5d \n", "j", j)
}

func relational() {

	// Relational operators
	fmt.Println("Relational operators")
	fmt.Printf("%-10s : %5d \n", "i:", i)
	fmt.Printf("%-10s : %5d \n", "j", j)
	fmt.Printf("%-10s : %t \n", "i == j", i == j)
	fmt.Printf("%-10s : %t \n", "i < j", i < j)
	fmt.Printf("%-10s : %t \n", "i <= j", i <= j)
	fmt.Printf("%-10s : %t \n", "i > j", i > j)
	fmt.Printf("%-10s : %t \n", "i >= j", i >= j)
	fmt.Printf("%-10s : %t \n", "i != j", i != j)
}

func bitwiseLogical() {
	// Bitwise logical operators
	//fmt.Println(strconv.FormatInt(int64(i), 2))
	//fmt.Println(strconv.FormatInt(int64(j), 2))
	fmt.Println("\nBitwise logical operators")
	fmt.Printf("%-10s : %064b \n", "i:", i)
	fmt.Printf("%-10s : %064b \n", "j", j)
	fmt.Printf("%-10s : %d \n", "i&j", i&j)
	fmt.Printf("%-10s : %d \n", "i|j", i|j)
	fmt.Printf("%-10s : %d \n", "i^j", i^j)
}

func shift() {
	// Shift operators
	//fmt.Println(strconv.FormatInt(int64(i), 2))
	i = 32
	fmt.Printf("%-10s : %d \n", "i:", i)
	fmt.Printf("%-10s : %d \n", "i>>2", i>>2)
	fmt.Printf("%-10s : %d \n", "i<<4", i<<4)
}
