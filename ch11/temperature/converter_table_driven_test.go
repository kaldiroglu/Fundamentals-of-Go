package converter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestC2FTableDriven(t *testing.T) {
	//Arrange
	var data = []struct {
		inputC    float64
		expectedF float64
	}{
		{-40.0, -40.0},
		{0.0, 32.0},
		{10.0, 50.0},
		{20.0, 68.0},
		{32.0, 89.6},
		{38.0, 100.4},
	}

	// Act & Assert
	for _, test := range data {
		actualF, _ := C2F(test.inputC)
		if test.expectedF != actualF {
			t.Logf("Expecting %f but got %f", test.expectedF, actualF)
		}
	}
}

func TestF2CTableDriven(t *testing.T) {
	//Arrange
	var data = []struct {
		inputF    float64
		expectedC float64
	}{
		{-40.0, -40.0},
		{20.0, -7.0},
		{32.0, 0.0},
		{40.0, 4.0},
		{80.0, 27.0},
		{100.4, 38.0},
	}

	// Act & Assert
	for _, test := range data {
		actualC, _ := F2C(test.inputF)
		assert.InDelta(t, test.expectedC, actualC, 0.6, "Should be in delta.")
	}
}
