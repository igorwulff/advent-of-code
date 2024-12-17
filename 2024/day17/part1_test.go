package main

import (
	"testing"
)

func TestPart1Test1(t *testing.T) {
	input := `Register A: 10
Register B: 0
Register C: 0

Program: 5,0,5,1,5,4`

	expected := "0,1,2"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestPart1Test2(t *testing.T) {
	input := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

	expected := "4,2,5,6,7,7,7,7,3,1,0"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestPart1Sample(t *testing.T) {
	input := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

	expected := "4,6,3,5,6,3,5,2,1,0"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
