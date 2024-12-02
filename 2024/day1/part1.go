package main

import (
	"fmt"
	"sort"

	"github.com/igorwulff/advent-of-code/2024/day1/shared"
	"github.com/igorwulff/advent-of-code/utils"
)

// Exported function to be called by the main application
func Part1(input string) string {
	left, right := shared.ParseInput(input)
	sortAsc(&left)
	sortAsc(&right)

	return fmt.Sprint(getDist(left, right))
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
