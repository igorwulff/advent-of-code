package main

import (
	"testing"
)

func TestPart1Sample(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`

	expected := "140"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestPart1Sample2(t *testing.T) {
	input := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

	expected := "772"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestPart1Sample3(t *testing.T) {
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

	expected := "1930"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
