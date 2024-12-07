package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day7/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	equations := shared.ParseInput(input)

	sum := 0
	for _, e := range equations {
		if e.IsValid() {
			sum += e.Result
		}
	}

	return fmt.Sprint(sum)
}
