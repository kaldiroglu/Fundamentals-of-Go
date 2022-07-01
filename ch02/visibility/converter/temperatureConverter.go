// Package converter includes a function to convert between celsius, fahrenheit, and kelvin.
package converter

// Convert function makes the conversions between celsius, fahrenheit, and kelvin.
func Convert(from byte, to byte, degree float64) float64 {
	var degreeConverted float64 = 0
	if from == 'c' || from == 'C' {
		if to == 'c' || to == 'C' {
			degreeConverted = degree
		} else if to == 'f' {
			degreeConverted = convertFromCelsiusToFahrenheit(degree)
		} else if to == 'k' {
			degreeConverted = convertFromCelsiusToKelvin(degree)
		}
	} else if from == 'f' || from == 'F' {
		if to == 'f' || to == 'F' {
			degreeConverted = degree
		} else if to == 'c' {
			degreeConverted = convertFromFahrenheitToCelsius(degree)
		} else if to == 'k' {
			degreeConverted = convertFromFahrenheitToKelvin(degree)
		}
	} else if from == 'k' || from == 'K' {
		if to == 'k' || to == 'K' {
			degreeConverted = degree
		} else if to == 'c' {
			degreeConverted = convertFromKelvinToCelsius(degree)
		} else if to == 'f' {
			degreeConverted = convertFromKelvinToFahrenheit(degree)
		}
	}

	return degreeConverted
}

func convertFromCelsiusToFahrenheit(degree float64) float64 {
	return (9 * degree / 5) + 32
}

func convertFromCelsiusToKelvin(degree float64) float64 {
	return degree + 273
}

func convertFromFahrenheitToCelsius(degree float64) float64 {
	return 5 * (degree - 32) / 9
}

func convertFromFahrenheitToKelvin(degree float64) float64 {
	return convertFromFahrenheitToCelsius(degree) + 273
}

func convertFromKelvinToCelsius(degree float64) float64 {
	return degree - 273
}

func convertFromKelvinToFahrenheit(degree float64) float64 {
	return (9 * (degree - 273) / 5) + 32.0
}
