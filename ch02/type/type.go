package main

import "fmt"

type yazi string

func main() {

	var s1 yazi = "I love Go!"
	fmt.Printf("%s\n", s1)

	var s2 string = "I love Go!"
	fmt.Printf("%s\n", s2)

	//fmt.Println(s1 == s2) //Nop! That's related to assignability
	fmt.Printf("The same: %t\n", string(s1) == s2)

	fmt.Printf("The same: %t\n", s1 == "I love Go!")
	fmt.Printf("The same: %t\n", "I love Go!" == s2)

	type sayi int
	var i sayi = 3
	fmt.Printf("%d\n", i)

	type text string
	var t text = "I love Go!"
	//fmt.Println(s == t) // That's related to assignability

	//var ss string = "I love Go!"
	//fmt.Println(s1 == ss)

	var s3 string = string(t)
	fmt.Println(s3)
	var s4 string = string(t)
	fmt.Println(s4)
	fmt.Println(s3 == s4)
}
