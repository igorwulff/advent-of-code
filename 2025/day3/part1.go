package main

import (
	"fmt"
	"strconv"

	"github.com/igorwulff/advent-of-code/2025/day3/shared"
)

func Part1(input string) string {
	sum := 0

	data := shared.ParseInput(input)

	for _, line := range data {
		first := 0
		second := 0
		idx := 0

		// first find highest value... and position
		for k, i := range line {
			if k+1 == len(line) {
				continue
			}

			val, _ := strconv.Atoi(string(i))
			if val > first {
				first = val
				idx = k
			}
		}

		for k, i := range line {
			if k <= idx {
				continue
			}

			val, _ := strconv.Atoi(string(i))
			if val > second {
				second = val
			}
		}

		sum += (first * 10) + second
	}

	return fmt.Sprint(sum)
}
