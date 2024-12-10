package main

import (
	"testing"
)

func TestPart2Sample(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	expected := "81"
	output := Part2(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
