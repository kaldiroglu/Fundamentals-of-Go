package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

type Singleton struct {
	name string
}

var singleton *Singleton
var instanceCount int

// NewSingleton has a problem: First if statement needs to be checked under lock only before the singleton is created.
// After it is created there is no need to be checked with lock, which may cause performance problems!
func NewSingleton(s string) (*Singleton, error) {
	var err error
	mu.Lock()
	if singleton != nil {
		err = errors.New("it is a singleton")
	} else {
		time.Sleep(1 * time.Millisecond)
		singleton = new(Singleton)
		singleton.name = s
		instanceCount++
		fmt.Println("Created :", instanceCount)
	}
	mu.Unlock()
	return singleton, err
}

func main() {

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("i%d", i)
		go NewSingleton(name)
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("\nCount is %d\n", instanceCount)
}
