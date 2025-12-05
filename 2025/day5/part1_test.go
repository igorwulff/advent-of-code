package main

import (
	"testing"
)

func TestPart1Sample(t *testing.T) {
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

	expected := "3"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
