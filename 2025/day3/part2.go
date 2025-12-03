package main

import (
	"fmt"
	"strconv"

	"github.com/igorwulff/advent-of-code/2025/day3/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	sum := 0

	// 12 batteries
	data := shared.ParseInput(input)

	for _, line := range data {
		val := ""
		start := 0

		for o := 12; o > 0; o-- {
			var highest int
			highest, start = FindHighestValue(start, line, o)

			val += strconv.Itoa(highest)
		}

		number, _ := strconv.Atoi(val)
		sum += number
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
