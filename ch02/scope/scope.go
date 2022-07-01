package main

import (
	"fmt"
	"os"
)

var s string = "I love Go!"

func main() {
	fmt.Printf("%s\n", s)

	var s string = "I really love Go!"
	fmt.Printf("%s\n", s)

	//var s string = "I said you, I really love Go!"
	//fmt.Printf("%s\n", s)
	//x := "hellooww"
	for i, s := range s { // That's so strange!
		s := s
		fmt.Printf("%2d:  %c\n", i, s)
	}
	fmt.Printf("%s\n", s)

	fmt.Println("*****************")

	length := len(s)
	fmt.Printf("Length is %d\n", length)
	if length := f(); length > length { // Change > to == and see the difference
		fmt.Printf("In if %d\n", length)
	} else {
		fmt.Printf("In else %d\n", length)
	}
	fmt.Printf("%d\n", length)

	fmt.Println("*****************")
}

func f() int {
	return 20
}

// GOPL p. 48
func g1() error {
	//if f, err := os.Open("scope.go"); err != nil { // compile error: unused: f
	//	return err
	//}
	//f.Stat() // compile error: undefined f
	//f.Close()    // compile error: undefined f
	return nil
}

func g2() error {
	f, err := os.Open("/Users/akin/Development/Go/Projects/Fundamentals of Go/ch02/scope/scope.go")
	if err != nil {
		return err
	}
	f.Stat()  // compile error: undefined f
	f.Close() // compile error: undefined f
	return nil
}
