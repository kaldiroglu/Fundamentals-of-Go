package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {

	start := time.Now()
	for i := 0; i < 1_000_000_000; i++ {
		computeWithLock(float64(i), float64(10*i))
	}
	duration := time.Since(start)
	fmt.Printf("Time is : %v", duration)
}

func compute(i float64, j float64) {
	math.Sqrt(i * j)
}
func computeWithLock(i float64, j float64) {
	mu.Lock()
	math.Sqrt(i * j)
	mu.Unlock()
}
