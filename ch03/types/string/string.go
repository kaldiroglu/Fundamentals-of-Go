package main

import (
	"fmt"
	"strings"
)

func main() {
	//basic()
	stringLiterals()
}

func basic() {
	s := "Selam, naber?"
	fmt.Printf("%s\n", s)
	fmt.Println(len(s))
	fmt.Println(s[1], s[9])
	fmt.Printf("%c %c\n", s[1], s[9])
	//fmt.Println(s[13])

	fmt.Println(s[1:3])
	fmt.Println(s[:])

	s1 := "selam, naber?"
	fmt.Println(s < s1)

	t := s
	s += " :)"
	fmt.Printf("%s \n%s\n", s, t)

	//s[0] = 'x'

	string1 := "string1"
	string2 := string1

	string3 := string1 + " " + string2
	fmt.Printf("string: \n%s", string3)
}

func stringLiterals() {
	fmt.Printf("%s\n", "String Literals")
	fmt.Println("I love Go!")
	poemString := "\n          Leylim Ley           \n" +
		"\tDöndüm daldan düşen kuru yaprağa\n" +
		"Seher yeli dağıt beni kır beni\n" +
		"\tGötür tozlarımı burdan uzağa\n" +
		"Yarin çıplak ayağına sür beni\n" +
		"...\n" +
		"Sabahattin Ali"

	fmt.Printf("%s\n", poemString)

	poemRawLiteralString := `

		         Leylim Ley
		Döndüm daldan düşen kuru yaprağa
				Seher yeli dağıt beni kır beni
		Götür tozlarımı burdan uzağa
				Yarin çıplak ayağına sür beni

		Sabahattin Ali`

	fmt.Printf("%s\n", poemRawLiteralString)
}

func manipulateStrings() {
	string1 := "string1"
	string2 := string1
	string3 := strings.ReplaceAll(string1, "1", "X")
	fmt.Printf("\n%s %s %s", string1, string2, string3)

	string4 := strings.ToUpper(string3)
	fmt.Printf("\n%s %s", string3, string4)

	string5 := string1 + " " + string2 + " " + string3
	fmt.Printf("\n%s", string5)
}
