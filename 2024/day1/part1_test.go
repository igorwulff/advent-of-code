package main

import (
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	input := []int{3, 4, 2, 1, 3, 3}

	expected := []int{1, 2, 3, 3, 3, 4}

	sortAsc(&input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}
