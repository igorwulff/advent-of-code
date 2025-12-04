package main

import (
	"testing"
)

func TestPart1Sample(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	expected := "13"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
