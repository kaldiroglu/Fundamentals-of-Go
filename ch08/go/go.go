package main

import (
	"fmt"
	"time"
)

var k int
var b bool = true

func main() {
	go1()
	//go2()

	time.Sleep(5 * time.Second)
	fmt.Println("Quiting!")
}

func go1() {
	go count() // Remove go and see its effect

	for i := 0; i < 10; i++ {
		fmt.Printf("%-2d: %-10d\n", i, k)
		time.Sleep(1 * time.Second)
	}
	b = false
}

func count() {
	for b {
		k++
	}
}

func go2() {
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Printf("Ali\n")
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Printf("Veli\n")
		}
	}()
}
