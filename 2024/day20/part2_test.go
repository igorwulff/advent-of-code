package main

import (
	"testing"

	"github.com/igorwulff/advent-of-code/2024/day20/shared"
)

func TestPart2Sample(t *testing.T) {
	shared.PicoSeconds = 50

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

	expected := "285"
	output := Part2(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
