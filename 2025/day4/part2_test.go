package main

import (
	"testing"
)

func TestPart2Sample(t *testing.T) {
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

	expected := "43"
	output := Part2(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
