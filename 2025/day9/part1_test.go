package main

import (
	"testing"
)

type Corner struct {
	X, Y int
}

type Distance struct {
	A    *Corner
	B    *Corner
	Dist int
}

func TestPart1Sample(t *testing.T) {
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

	expected := "50"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
