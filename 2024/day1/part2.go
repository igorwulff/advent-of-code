package main

import (
	"github.com/igorwulff/advent-of-code/2024/day1/shared"

	"fmt"
)

// Exported function to be called by the main application
func Part2(input string) string {
	left, right := shared.ParseInput(input)

	m := findOccurences(left, right)

	output := 0
	for _, l := range left {
		output += l * m[l]

	}

	return fmt.Sprint(output)
}

func findOccurences(left, right []int) map[int]int {
	m := make(map[int]int)

	for _, l := range left {
		i := 0

		for _, r := range right {
			if l == r {
				i++
			}
		}

		m[l] = i
	}

	return m
}
