package main

import (
	"testing"
)

func TestPart1Sample(t *testing.T) {
	input := `r, wr, b, g, bwu, rb, gb, br
brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

	expected := "6"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
