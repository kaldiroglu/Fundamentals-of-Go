package main

import (
	"fmt"
)

type Singleton struct {
	name          string
	instanceCount int
}

var singleton *Singleton

func NewSingleton(s string) *Singleton {
	singleton = new(Singleton)
	singleton.instanceCount++
	singleton.name = s
	return singleton
}

func main() {

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("i%d", i)
		s := NewSingleton(name)
		fmt.Println(s.name)
		fmt.Println(s.instanceCount)
	}
}
