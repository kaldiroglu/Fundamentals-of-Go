package main

import (
	"fmt"
	"time"
)

func main() {
	//createChannel()

	//var ch chan string = make(chan string)
	//go sendData(ch)
	//ch <- "Hey, what's up?"

	sendReceiveData()

	time.Sleep(100 * time.Millisecond)
}

func createChannel() {
	var ch chan string
	ch = make(chan string)
	fmt.Printf("%T & %#v", ch, ch)
}

func sendData(ch chan string) {
	fmt.Println("Data in channel is: ", <-ch)
}

func sendReceiveData() {
	var channel chan int
	channel = make(chan int)

	go send(channel)
	time.Sleep(10000 * time.Millisecond)
	go receive(channel)
}

func send(sender chan int) {
	fmt.Println("Sending")
	for i := 0; i < 100; i++ {
		fmt.Println("Data sent: ", i)
		sender <- i
	}
}

func receive(receiver chan int) {
	fmt.Println("Receiving")
	for i := 0; i < 100; i++ {
		fmt.Println("Data received: ", <-receiver)
	}
}
