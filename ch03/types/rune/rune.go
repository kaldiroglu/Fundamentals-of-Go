package main

import (
	"fmt"
)

func main() {
	fmt.Println('a')
	fmt.Println('\n')
	fmt.Println('\u000a')
	fmt.Printf("%c %c\n", 'n', '\a')

	var c1 rune = 'x'
	fmt.Printf("%c\n", c1)
	fmt.Printf("%T\n", c1)

	var c2 rune = '\u011f'
	fmt.Printf("Char: %c\n", c2)
	fmt.Printf("%T\n", c2)

	var c3 byte = 'g' // make it ğ
	fmt.Printf("%c\n", c3)
	fmt.Printf("%T\n", c3)

	var c4 rune = rune(c3)
	fmt.Printf("%c\n", c4)

	nl := '\n'
	fmt.Printf("%c", nl)
}
