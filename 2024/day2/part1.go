package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day2/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	levels := shared.ParseInput(input)

	count := 0
	for _, level := range levels {
		if level.IsSafe() {
			count++
		}
	}

	return fmt.Sprint(count)
}
