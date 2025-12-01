package main

import (
	"testing"
)

func TestPart2Sample(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	expected := "6"
	output := Part2(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
