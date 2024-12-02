package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/igorwulff/advent-of-code/utils"
)

// Exported function to be called by the main application
func Part1(input string) string {
	left, right := parseInput(input)
	sortAsc(&left)
	sortAsc(&right)

	return fmt.Sprint(getDist(left, right))
}

func parseInput(input string) ([]int, []int) {
	left := make([]int, 0)
	right := make([]int, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		var l, r int
		fmt.Sscanf(line, "%d %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func sortAsc(input *[]int) {
	sort.Slice(*input, func(i, j int) bool {
		return (*input)[i] < (*input)[j]
	})
}

func getDist(left []int, right []int) int {
	sum := 0

	for i, l := range left {
		r := right[i]
		sum += utils.AbsInt(r - l)
	}

	return sum
}
