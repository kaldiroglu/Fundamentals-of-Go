package main

import "fmt"

type Person struct {
	name string
}

func (p Person) greet() string {
	return "Selam from " + p.name + " :)"
}

func (p Person) greetSomebody(q Person) string {
	return "Selam from " + p.name + " to " + q.name + " :)"
}

func main() {
	var greeter Person = Person{"Nihal"}
	greeting := greeter.greet()
	fmt.Printf("%s\n", greeting)

	var myFriend Person = Person{"Mehmet"}
	greeting = greeter.greetSomebody(myFriend)
	fmt.Printf("%s\n", greeting)
}
