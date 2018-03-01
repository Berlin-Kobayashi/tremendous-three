package main

import (
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	a, b := Coordinates{0, 0}, Coordinates{1, 2}
	d := CalculateDistance(a, b)

	if d != 3 {
		t.Error("Math no good")
	}

	a, b = Coordinates{5, 6}, Coordinates{1, 2}
	d = CalculateDistance(a, b)

	if d != 8 {
		t.Error("Math no good")
	}
}
