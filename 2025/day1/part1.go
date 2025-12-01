package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2025/day1/shared"
)

func Part1(input string) string {
	dirs, steps := shared.ParseInput(input)
	counter := make(map[int]int, len(dirs))

	pos := 50

	for i, dir := range dirs {
		step := steps[i]

		switch dir {
		case "L":
			pos -= step
			pos = pos % 100
			if pos < 0 {
				pos += 100
			}
		case "R":
			pos += step
			pos = pos % 100
		}

		counter[pos]++
	}

	maxCount := 0
	for _, v := range counter {
		if v > maxCount {
			maxCount = v
		}
	}

	return fmt.Sprint(maxCount)
}
