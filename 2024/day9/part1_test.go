package main

import (
	"testing"
)

func TestPart1Sample(t *testing.T) {
	input := "2333133121414131402"
	expected := "1928"
	output := Part1(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
