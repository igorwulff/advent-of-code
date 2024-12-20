package main

import (
	"testing"

	"github.com/igorwulff/advent-of-code/2024/day20/shared"
)

func TestPart1Sample(t *testing.T) {
	shared.PicoSeconds = 0

	input := `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

	expected := "44"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
