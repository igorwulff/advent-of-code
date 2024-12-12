package main

import (
	"testing"
)

func TestPart2Sample(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`

	expected := "80"
	output := Part2(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestPart2Sample2(t *testing.T) {
	input := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

	expected := "436"
	output := Part2(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestPart2Sample3(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

	expected := "1206"
	output := Part2(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
