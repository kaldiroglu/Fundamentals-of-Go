package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var instanceCount int
var mu sync.Mutex

type Singleton struct {
	name string
}

var singleton *Singleton

// NewSingleton applies "double-checked locked pattern"
func NewSingleton(s string) (*Singleton, error) {
	var err error
	if singleton == nil {
		fmt.Println("Outside: Already created!")
		err = errors.New("it is a singleton")
	} else {
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		if instanceCount == 1 {
			fmt.Println("Inside: Already created!")
			err = errors.New("it is a singleton")
		} else {
			singleton = new(Singleton)
			singleton.name = s
			instanceCount++
			fmt.Println("Created :", instanceCount)
		}
		mu.Unlock()
	}
	return singleton, err
}

func main() {

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("i%d", i)
		go NewSingleton(name)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("Count is %d\n", instanceCount)
}

// https://turtledev.in/posts/go-concurrency-patterns-double-checked-locking/
