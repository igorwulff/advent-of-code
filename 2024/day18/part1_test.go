package main

import (
	"testing"

	"github.com/igorwulff/advent-of-code/2024/day18/shared"
)

func TestPart1Sample(t *testing.T) {
	input := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

	expected := "22"

	shared.CorruptedBytes = 12
	shared.Width = 7
	shared.Height = 7

	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
