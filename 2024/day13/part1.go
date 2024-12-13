package main

import (
	"fmt"
	"sort"

	"github.com/igorwulff/advent-of-code/2024/day13/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	claws := shared.ParseInput(input)

	sum := 0

	for _, claw := range claws {
		matches := claw.Matches()
		if len(matches) == 0 {
			continue
		}

		sort.Ints(matches)
		sum += matches[0]
	}

	return fmt.Sprint(sum)
}
