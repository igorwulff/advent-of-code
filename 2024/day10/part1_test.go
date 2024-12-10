package main

import (
	"testing"
)

/*func TestPart1Simple(t *testing.T) {
	input := `0123
1234
8765
9876`

	expected := "1"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestPart1Simple2(t *testing.T) {
	input := `2220222
2221222
2222222
6543456
7224227
8225678
9298729`

	expected := "3"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}*/

func TestPart1Sample(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	expected := "36"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
