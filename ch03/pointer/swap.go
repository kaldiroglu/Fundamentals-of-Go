package main

import "fmt"

func main() {
	x := 5
	y := 7
	fmt.Printf("x = %d y = %d\n\n", x, y)

	x, y = swap1(x, y)
	fmt.Printf("x = %d y = %d\n", x, y)

	x = 5
	y = 7
	swap2(x, y)
	fmt.Printf("x = %d y = %d\n", x, y)

	x = 5
	y = 7
	swap3(&x, &y)
	fmt.Printf("x = %d y = %d\n", x, y)
}

func swap1(x, y int) (int, int) {
	temp := x
	x = y
	y = temp
	return x, y
}

func swap2(x, y int) {
	temp := x
	x = y
	y = temp
}

func swap3(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
