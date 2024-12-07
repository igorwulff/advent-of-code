package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day7/shared"
)

func Part2(input string) string {
	equations := shared.ParseInput(input)

	sum := 0
	for _, e := range equations {
		e.Concat = true
		if e.IsValid() {
			sum += e.Result
		}
	}

	return fmt.Sprint(sum)
}
