package main

import (
	"testing"
)

func TestPart2Sample(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`

	expected := "3121910778619"
	output := Part2(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
