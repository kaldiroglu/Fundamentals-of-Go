package main

import (
	"fmt"
	"reflect"
)

func main() {

	defaultValues()
	//defaultTypes()
}

func defaultValues() {
	var i int
	var f float64
	var b bool
	var by byte
	var r rune
	var s string
	var intArray [2]int
	var boolArray [2]bool
	var m map[string]string // = nil
	var p Person
	var ptr *int = nil

	fmt.Printf("%s\n", "Default Values")
	fmt.Print("Default value for a int: ")
	fmt.Println(i)
	fmt.Print("Default value for a float64: ")
	fmt.Println(f)
	fmt.Print("Default value for a bool: ")
	fmt.Println(b)
	fmt.Print("Default value for a byte: ")
	fmt.Println(by)
	fmt.Print("Default value for a rune: ")
	fmt.Println(r)
	fmt.Println(string(r))
	fmt.Print("Default value for a string: >" + s + "<")
	fmt.Print("Default Value for a array: ")
	fmt.Println(intArray)
	fmt.Print("Default Value for a array: ")
	fmt.Println(boolArray)
	fmt.Print("Default Value for a map: ")
	fmt.Println(m)
	fmt.Println(m == nil)
	fmt.Print("Default value for a struct: ")
	fmt.Println(p)
	fmt.Print("Default value for a pointer: ")
	fmt.Println(ptr)

	// nil can't be used for non-reference types!
	//  fmt.Println(i == nil)
	//  fmt.Println(&i == nil)
	//	fmt.Println(s == nil)
	//	fmt.Println(p == nil)
}

func defaultTypes() {
	fmt.Printf("%-10T %s\n", 5, reflect.TypeOf(5))
	fmt.Printf("%-10T %s\n", 3.14, reflect.TypeOf(3.14))
	fmt.Printf("%-10T %s\n", 'a', reflect.TypeOf('a'))
	fmt.Printf("%-10T %s\n", "a", reflect.TypeOf("a"))
	fmt.Printf("%-10T %s\n", false, reflect.TypeOf(false))
}

type Person struct {
	id   int
	name string
}
