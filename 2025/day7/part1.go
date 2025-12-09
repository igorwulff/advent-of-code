package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2025/day7/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	data := shared.ParseInput(input)

	row := data.Row
	split := 0
	for {
		s, ready := data.Step(row)
		split += s
		if ready {
			break
		}

		row++
	}

	return fmt.Sprint(split)
}
