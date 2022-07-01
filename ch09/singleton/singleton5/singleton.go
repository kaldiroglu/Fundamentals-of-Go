package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var instanceCount int
var mu sync.Mutex
var rwmu sync.RWMutex

type Singleton struct {
	name string
}

var singleton *Singleton

// NewSingleton1 applies "double-checked locked pattern"
func NewSingleton1(s string) (*Singleton, error) {
	var err error
	if singleton != nil {
		fmt.Println("Outside: Already created!")
		err = errors.New("it is a singleton")
	} else {
		//time.Sleep(10 * time.Millisecond)
		mu.Lock()
		if singleton != nil {
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

// NewSingleton2 applies sync.RLock
func NewSingleton2(s string) (*Singleton, error) {
	var err error
	if singleton == nil {
		rwmu.RLock()
		if singleton == nil {
			mu.Lock()
			singleton = new(Singleton)
			singleton.name = s
			instanceCount++
			mu.Unlock()
			fmt.Println("Created :", instanceCount)
		} else {
			fmt.Println("Inside: Already created!")
			err = errors.New("it is a singleton")
		}
		rwmu.RUnlock()
	} else {
		fmt.Println("Outside: Already created!")
		err = errors.New("it is a singleton")
	}
	return singleton, err
}

func main() {

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("i%d", i)
		go NewSingleton2(name)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("Count is %d\n", instanceCount)
}

// https://turtledev.in/posts/go-concurrency-patterns-double-checked-locking/
