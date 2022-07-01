package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	basics()
	//fromStringToNumeric()
	//fromNumericToString()
}

func basics() {
	var fp float64 = 3.14
	fmt.Println(fp)
	var i int
	i = int(fp)
	fmt.Println(i)

	//i = int(6.28) // But, cannot convert an expression of the type 'float64' to the type 'int'
	i = int(math.Floor(6.28))
	fmt.Println(i)

	var j int64 = 123_456_789_012
	fmt.Println(j)
	var k int32 = int32(j)
	fmt.Println(k)

	type yazi string
	var y yazi = "I like Go!"
	type text string
	var t text = "I love Go!"
	var s1 string = string(y)
	var s2 string = string(t)
	fmt.Println(s1, s2)

	//fmt.Println(y == t) //Nop! That's related to assignability
	fmt.Println(string(y) == string(t))
}

func fromStringToNumeric() {
	println("\nfromStringToNumeric")
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("%f\n", f)
	i, _ := strconv.ParseInt("-19", 10, 64)
	fmt.Printf("%d\n", i)
	ui, _ := strconv.ParseUint("91", 10, 64)
	fmt.Printf("%d\n", ui)
}

func fromNumericToString() {
	println("\nfromNumericToString")
	s := strconv.FormatFloat(314159265358979323846.26433832795028841971693993751, 'E', 20, 64)
	fmt.Printf("%s\n", s)
	s = strconv.FormatInt(1234567890, 10)
	fmt.Printf("%s\n", s)
	s = strconv.FormatUint(91, 10)
	fmt.Printf("%s\n", s)
}

func convertBoolAndString() {
	var b bool = false
	//boolString := string(b)
	var boolString string = strconv.FormatBool(b)
	fmt.Printf("The same: %s\n", boolString)

	boolString = "true"
	b, _ = strconv.ParseBool(boolString)
	fmt.Printf("The same: %t\n", b)
}
