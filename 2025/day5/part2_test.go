package main

import (
	"testing"
)

func TestPart2Sample(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	expected := "14"
	output := Part2(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
