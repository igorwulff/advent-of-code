package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day19/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	towels, designs := shared.ParseInput(input)

	valid := 0
	for _, design := range designs {
		valid += towels.CountWays(design)
	}

	return fmt.Sprint(valid)
}
