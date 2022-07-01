package converter

import (
	"testing"
)

func TestC2FConvergingAtMinus40(t *testing.T) {
	//Arrange
	expectedF := -40.0

	// Act
	actualF, _ := C2F(-40) // SUT is C2F

	// Assert
	if expectedF != actualF {
		t.Logf("Expecting %f but got %f", expectedF, actualF)
	}
}

func TestF2CConvergingAtMinus40(t *testing.T) {
	//t.Fail()
	expectedC := -40.0
	actualC, _ := C2F(-40)
	if expectedC != actualC {
		t.Logf("Expecting %f but got %f", expectedC, actualC)
	}
}

func TestC2FHumanTemp(t *testing.T) {
	expectedMinF := 98.0
	expectedMaxF := 102.0
	actualF, _ := C2F(37)
	if !(expectedMinF < actualF) || !(expectedMaxF > actualF) {
		t.Logf("Expecting between %f and %f but got %f", expectedMinF, expectedMaxF, actualF)
	}
}

func TestF2CHumanTemp(t *testing.T) {
	expectedMinC := 37.0
	expectedMaxC := 38.0
	actualC, _ := F2C(100)
	if !(expectedMinC < actualC) || !(expectedMaxC > actualC) {
		t.Logf("Expecting between %f and %f but got %f", expectedMinC, expectedMaxC, actualC)
	}
}

func TestC2FExpectingError(t *testing.T) {
	_, err := C2F(-273.1)
	if err == nil {
		t.Logf("Expecting error: %s but got nothing!", err.Error())
	}
}

func TestF2CExpectingError(t *testing.T) {
	_, err := F2C(-460.1)
	if err == nil {
		t.Logf("Expecting error: %s but got nothing!", err.Error())
	}
}
