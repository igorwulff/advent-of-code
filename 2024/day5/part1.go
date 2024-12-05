package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day5/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	p := shared.ParseInput(input)

	sum := p.GetSorted(true)

	return fmt.Sprint(sum)
}
