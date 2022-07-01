package main

import (
	"errors"
	"fmt"
	"time"
)

type Singleton struct {
	name string
}

var singleton *Singleton
var instanceCount int

// NewSingleton is also broken since when threads accumulate at the beginning of the if statement they all check whether instanceCount is 1.
func NewSingleton(s string) (*Singleton, error) {
	var err error
	if singleton != nil {
		err = errors.New("it is a singleton")
	} else {
		//time.Sleep(1 * time.Millisecond)
		singleton = new(Singleton)
		singleton.name = s
		instanceCount++
		fmt.Println("Created :", instanceCount)
	}

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
