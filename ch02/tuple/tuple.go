package main

import (
	"fmt"
	"strconv"
)

func main() {
	calculateFibonacci(30)
}

func calculateFibonacci(n int) {
	print("Calculating first " + strconv.Itoa(n) + " Fibonacci numbers.\n")
	f0, f1 := 0, 1
	fmt.Printf("%3d %10d\n", 1, f0)
	fmt.Printf("%3d %10d\n", 2, f1)
	for i := 0; i <= n-3; i++ {
		f2 := f0 + f1
		fmt.Printf("%3d %10d %10f\n", i+3, f2, float64(f1)/float64(f2))
		f0, f1 = f1, f2
	}
}
