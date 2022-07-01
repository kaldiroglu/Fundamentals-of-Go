package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	TL
	OTHER = -1
)

// GOPL p. 77
type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func main() {
	fmt.Println(USD, EUR, GBP, TL, OTHER)
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)

	//usd := 1
	//processCurrency1(usd)
	//euro := 2
	//processCurrency1(euro)

	//processCurrency2(USD)
	//processCurrency2(TL)

	//i := 12
	//processCurrency2(i)
	processCurrency2(12)
}

func processCurrency1(currency int) {
	if currency == 0 {
		fmt.Println("It is USD")
	} else if currency == 1 {
		fmt.Println("It is EUR")
	} else if currency == 2 {
		fmt.Println("It is GBP")
	} else if currency == 3 {
		fmt.Println("It is TL")
	} else if currency == -1 {
		fmt.Println("It is OTHER")
	} else {
		fmt.Println("It is unknown currency!")
	}
	fmt.Println(currency)
}

func processCurrency2(currency Currency) {
	if currency == USD {
		fmt.Println("It is USD")
	} else if currency == EUR {
		fmt.Println("It is EUR")
	} else if currency == GBP {
		fmt.Println("It is GBP")
	} else if currency == TL {
		fmt.Println("It is TL")
	} else if currency == -1 {
		fmt.Println("It is OTHER")
	} else {
		fmt.Println("It is unknown currency!")
	}
	fmt.Println(currency)
}
