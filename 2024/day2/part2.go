package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day2/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	levels := shared.ParseInput(input)

	count := 0
	for _, level := range levels {
		if level.IsSafe(false) || level.IsSafe(true) {
			count++
		}
	}

	return fmt.Sprint(count)
}
