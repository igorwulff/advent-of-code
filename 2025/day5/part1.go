package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2025/day5/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	ingr := shared.ParseInput(input)

	count := 0
	for _, v := range ingr.Avail {
		if ingr.CheckIfExists(v) {
			count++
		}
	}

	return fmt.Sprint(count)
}
