package main

import (
	"testing"
)

func TestPart1Simple(t *testing.T) {
	input := `..........
..........
..........
....a.....
..........
.....a....
..........
..........
..........
..........`

	expected := "2"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestPart1Simple2(t *testing.T) {
	input := `..........
..........
..........
....a.....
........a.
.....a....
..........
..........
..........
..........`

	expected := "4"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}

}

func TestPart1Sample(t *testing.T) {
	input := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	expected := "14"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
