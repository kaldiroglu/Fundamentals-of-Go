package main

import (
	"fmt"
	"sync"
	"time"
)

// https://go.dev/ref/mem
// https://medium.com/@genchilu/the-difference-between-java-and-golang-in-writing-concurrent-code-to-access-shared-variable-9e59aa7d789d
var mu sync.Mutex

func main() {
	case1()
	//case2()
}

func case1() {
	var x, y int
	go func() {
		x = 1
		fmt.Println("y:", y, " ") // A2
	}()
	go func() {
		y = 1
		fmt.Print("x:", x, " ") // B2
	}()

	//fmt.Println(x, y)
	time.Sleep(10 * time.Millisecond)
	fmt.Println(x, y)
}

func case2() {
	var flag bool = true

	go func() {
		var i int
		for flag {
			fmt.Println(i)
			i++
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		mu.Lock()
		flag = false
		mu.Unlock()
	}()

	time.Sleep(10 * time.Second)
}
