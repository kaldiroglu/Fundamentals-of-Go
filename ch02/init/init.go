package main

import "fmt"

var j int = i + 1

func createB() bool {
	fmt.Println("In createB")
	return true
}

var i int = createI()

func main() {
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(b)
}

func createI() int {
	fmt.Println("In createI")
	//fmt.Printf("%d %d \n", i, j)
	return 5
}

func init() {
	i := 7
	fmt.Println("In init I")
	fmt.Printf("%d %d \n", i, j)
}

func init() {
	fmt.Println("In init II")
	print(b)
	fmt.Printf("%d %d \n", i, j)
}

var b bool //= createB()
