package main

import (
	"fmt"
	"sort"

	"github.com/igorwulff/advent-of-code/2024/day13/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	claws := shared.ParseInput(input)

	sum := 0

	for _, claw := range claws {
		claw.Prize.X += 1000000000
		claw.Prize.Y += 1000000000

		matches := claw.Matches()
		if len(matches) == 0 {
			continue
		}

		sort.Ints(matches)
		sum += matches[0]
	}

	return fmt.Sprint(sum)
}
