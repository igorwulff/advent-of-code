package main

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/igorwulff/advent-of-code/2024/day5/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	printer := shared.ParseInput(input)

	sum := 0

	for el := printer.Updates.Front(); el != nil; el = el.Next() {
		sorted := make([]int, len(el.Value))
		copy(sorted, el.Value)

		sort.Slice(sorted, func(i, j int) bool {
			r := printer.Rules[el.Value[j]]

			for _, v := range r {
				if v == el.Value[i] {
					return false
				}
			}

			return true
		})

		if reflect.DeepEqual(sorted, el.Value) {
			sum += el.Value[(len(el.Value)-1)/2]
		}
	}

	return fmt.Sprint(sum)
}
