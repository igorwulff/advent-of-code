package main

import (
	"testing"
)

func TestPart1Simple(t *testing.T) {
	input := `p=2,4 v=2,-3`
	width = 11
	height = 7

	expected := "0"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestPart1Sample(t *testing.T) {
	input := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

	width = 11
	height = 7

	expected := "12"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
