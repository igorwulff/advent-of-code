package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/igorwulff/advent-of-code/2025/day2/shared"
)

func Part2(input string) string {
	ids := shared.ParseInput(input)
	sum := 0

	for _, id := range ids {
		for i := id.Start; i <= id.End; i++ {
			value := strconv.Itoa(i)
			if !validateComplex(value) {
				sum += i
			}
		}
	}

	return fmt.Sprint(sum)
}

func validateComplex(input string) bool {
	for i := 1; i <= len(input)/2; i++ {
		// Check if number of characters can be divided by its total.
		// "123" can only be divided by 1 or by 3, not by 2.
		if len(input)%i > 0 {
			continue
		}

		part := input[:i]
		count := len(input) / i

		if strings.Repeat(part, count) == input {
			return false
		}
	}

	return true
}
