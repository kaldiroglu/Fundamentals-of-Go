// Run: go mod edit -replace  temperature/converter=./converter

package main

import (
	"fmt"
	"temperature/converter"
	tmp "temperature/converter"
)

func main() {

	cDegree := 36.7
	fDegree := converter.Convert('c', 'f', cDegree)
	fmt.Printf("%f celcius is %f fahrenheit.\n", cDegree, fDegree)

	cDegree = tmp.Convert('f', 'c', -40)
	fmt.Printf("%d celcius is %f fahrenheit.", -40, cDegree)

	//fDegree = converter.convertFromCelsiusToFahrenheit(fDegree)
}
