package main

import (
	"testing"
)

func TestPart2Sample(t *testing.T) {
	input := "2333133121414131402"
	expected := "2858"
	output := Part2(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
