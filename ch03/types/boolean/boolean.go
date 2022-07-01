package main

import "fmt"

func main() {

	b1 := true
	b2 := false
	fmt.Printf("%t and %t\n", b1, b2)

	b3 := 2 == 2
	b4 := 3 == 2
	fmt.Printf("%t and %t\n", b3, b4)

	if b3 && b4 {
		fmt.Println("Both 2 == 2 and 3 == 2 is true!")
	} else {
		fmt.Println("Either 2 == 2 or 3 == 2 or both are false!")
	}

	if b3 || b4 {
		fmt.Println("Either 2 == 2 or 3 == 2 or both are true!")
	} else {
		fmt.Println("Neither n2 == 2 or 3 == 2 is true!")
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", i)
	}

	//var i int = 0
	//var j int = 1

	//b1 = i       // Can't do that!
	//b1 = bool(i) // Can't do that!
}
