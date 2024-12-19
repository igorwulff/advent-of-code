package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day19/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	towels, designs := shared.ParseInput(input)

	valid := 0
	for _, design := range designs {
		if towels.Match(design) {
			valid++
		}
	}

	return fmt.Sprint(valid)
}
