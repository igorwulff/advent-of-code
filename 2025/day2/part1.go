package main

import (
	"fmt"
	"strconv"

	"github.com/igorwulff/advent-of-code/2025/day2/shared"
)

func Part1(input string) string {
	ids := shared.ParseInput(input)
	sum := 0

	for _, id := range ids {
		for i := id.Start; i <= id.End; i++ {
			value := strconv.Itoa(i)
			if !validate(value) {
				sum += i
			}
		}
	}

	return fmt.Sprint(sum)
}

func validate(input string) bool {
	l := input[:len(input)/2]
	r := input[len(input)/2:]

	return l != r
}
