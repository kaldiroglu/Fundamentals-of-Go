package main

import "fmt"

func main() {
	//intLiterals()
	//fpLiterals()
	//runeLiterals()
	//stringLiterals()
}

func intLiterals() {
	fmt.Printf("%s\n", "Decimal Literals")
	fmt.Println(5)
	fmt.Printf("%d\n", 15)

	fmt.Printf("\n%s\n", "Binary Literals")
	fmt.Println(0b11)
	fmt.Printf("%d\n", 0b1111) // Lower-case b
	fmt.Printf("%d\n", 0b1000) // Upper-case B

	fmt.Printf("\n%s\n", "Octal Literals")
	fmt.Println(0o77)
	fmt.Printf("%d\n", 0o1234) // Lower-case o
	fmt.Printf("%d\n", 0o1234) // Upper-case O

	fmt.Printf("\n%s\n", "Hexadecimal Literals")
	fmt.Println(0xAcaba)
	fmt.Printf("%d\n", 0xffff) // Lower-case x
	fmt.Printf("%d\n", 0xFFFF) // Upper-case X

	fmt.Printf("\n%s\n", "Literals with Underscore")
	fmt.Println(1_000_000_000)
	fmt.Printf("%d\n", 0b1001_1111)
	fmt.Printf("%d\n", 0xFFFF_FFFF)
}

func fpLiterals() {
	fmt.Printf("%s\n", "Floaintg-Point Literals")
	fmt.Println(0.5)
	fmt.Printf("%f\n", 0.5)

	fmt.Printf("%s\n", "FP with Exponent Literals")
	fmt.Println(5e-1)
	fmt.Printf("%f\n", 1.6e-19)
	fmt.Printf("%f\n", 6.022e23)

	fmt.Printf("%s\n", "Hexadecimal FP Literals")
	fmt.Println(0xAcaba)
	fmt.Printf("%d\n", 0xffff) // Lower-case x
	fmt.Printf("%d\n", 0xFFFF) // Upper-case X

	fmt.Printf("%s\n", "Literals with Underscore")
	fmt.Println(1_000_000_000.0)
	fmt.Printf("%f\n", 1.0_0e10)
	fmt.Printf("%f\n", 0xA.1p0)
	fmt.Printf("%f\n", 0xFFFF_FFFF.AAAAp-20)
}

func runeLiterals() {
	fmt.Printf("%s\n", "Rune Literals")
	fmt.Println('a')
	fmt.Println(string('a'))
	fmt.Printf("%d %c\n", 'A', 'A')
	fmt.Printf("%d %c\n", 'ğ', 'ğ')
	fmt.Printf("%d %c\n", 'ı', 'i')
}

func stringLiterals() {
	fmt.Printf("%s\n", "String Literals")
	fmt.Println("I love Go!")
	poemString := "\n          Leylim Ley           \n" +
		"Döndüm daldan düşen kuru yaprağa\n" +
		"Seher yeli dağıt beni kır beni\n" +
		"Götür tozlarımı burdan uzağa\n" +
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
