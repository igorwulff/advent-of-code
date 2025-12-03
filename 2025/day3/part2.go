package main

import (
	"fmt"
	"strconv"

	"github.com/igorwulff/advent-of-code/2025/day3/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	sum := 0

	data := shared.ParseInput(input)

	for _, line := range data {
		value := 0
		start := 0

		for offset := 12; offset > 0; offset-- {
			var highest int
			highest, start = FindHighestValue(start, line, offset)
			value = (value * 10) + highest
		}

		sum += value
	}

	return fmt.Sprint(sum)
}

func FindHighestValue(start int, line string, offset int) (highest int, idx int) {
	highest = 0
	idx = 0

	for i := start; i <= len(line)-offset; i++ {
		val, _ := strconv.Atoi(string(line[i]))
		if val > highest {
			highest = val
			idx = i
		}
	}

	return highest, idx + 1
}
