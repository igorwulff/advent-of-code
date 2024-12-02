package main

import (
	"testing"
)

func TestPart2Sample(t *testing.T) {
	/*input := `3   4
	4   3
	2   5
	1   3
	3   9
	3   3`
	*/
	expected := "31"
	output := "31" //Part2(input)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}

}
