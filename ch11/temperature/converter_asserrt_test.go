package converter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestC2FConvergingAtMinus40WithAssert(t *testing.T) {
	//Arrange
	expectedF := -40.0

	// Act
	actualF, _ := C2F(-40)

	// Assert
	//if expectedF != actualF {
	//	t.Logf("Expecting %f but got %f", expectedF, actualF)
	//}

	assert.Equal(t, expectedF, actualF, "Should be the same")
}

func TestF2CConvergingAtMinus40WithAssert(t *testing.T) {
	//t.Fail()
	expectedC := -40.0
	actualC, _ := C2F(-40)
	assert.Equal(t, expectedC, actualC, "Should be the same")
}

func TestC2FHumanTempWithAssert(t *testing.T) {
	expectedF := 98.0
	actualF, _ := C2F(37)
	//if !(expectedMinF < actualF) || !(expectedMaxF > actualF) {
	//	t.Logf("Expecting between %f and %f but got %f", expectedMinF, expectedMaxF, actualF)
	//}
	assert.InDelta(t, expectedF, actualF, 0.6, "Should be in delta.")
}

func TestF2CHumanTempWithAssert(t *testing.T) {
	expectedC := 37.0
	actualC, _ := F2C(98)
	assert.InDelta(t, expectedC, actualC, 0.6, "Should be in delta.")
}

func TestC2FExpectingErrorWithAssert(t *testing.T) {
	_, err := C2F(-273.1)
	//if err == nil {
	//	t.Logf("Expecting error: %s but got nothing!", err.Error())
	//}
	assert.Error(t, err, "Must return error.")
}

func TestF2CExpectingErrorWithAssert(t *testing.T) {
	_, err := F2C(-460.1)
	assert.Error(t, err, "Must return error.")
}
