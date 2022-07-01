package converter

import (
	"errors"
	"fmt"
)

// These are the constants representing the lowest possible degrees.
const (
	lowestC = -273.0
	lowestF = -460.0
)

// C2F converts from celsius to fahrenheit
func C2F(cDegree float64) (float64, error) {
	if cDegree < lowestC {
		return 0, errors.New("Invalid degree passed: " + fmt.Sprintf("%f", cDegree))
	} else {
		return (9 * cDegree / 5) + 32, nil
	}
}

// F2C converts from fahrenheit to celsius.
func F2C(fDegree float64) (float64, error) {
	if fDegree < lowestF {
		return 0, errors.New("invalid degree passed: " + fmt.Sprintf("%f", fDegree))
	}
	return (fDegree - 32) * 5 / 9, nil
}
