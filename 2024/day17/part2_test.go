package main

import (
	"testing"
)

func TestPart2Sample(t *testing.T) {
	input := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

	expected := "117440"
	output := Part2(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
