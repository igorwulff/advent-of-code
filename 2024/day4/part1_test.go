package main

import (
	"fmt"
	"testing"

	"github.com/igorwulff/advent-of-code/2024/day4/shared"
)

func TestPart1Sample(t *testing.T) {
	input := `....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX`

	expected := "18"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestDiagonalXMASPart1(t *testing.T) {
	input := `...S
..A.
.M..
X...`

	grid := shared.ParseInput(input, []string{"X", "S"})
	output := fmt.Sprint(grid.FindWords("XMAS"))

	expected := "0"
	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestDiagonalSAMXPart1(t *testing.T) {
	input := `S..S
.AA.
.MM.
X..X`

	grid := shared.ParseInput(input, []string{"X", "S"})
	output := fmt.Sprint(grid.FindWords("SAMX"))

	expected := "2"
	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
