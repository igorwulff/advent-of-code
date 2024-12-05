package shared

import (
	"reflect"
	"slices"
)

type Printer struct {
	Rules   map[int][]int
	Updates [][]int
}

func (p Printer) GetSorted(faulty bool) int {
	sum := 0

	for _, update := range p.Updates {
		sorted := make([]int, len(update))
		copy(sorted, update)

		slices.SortFunc(sorted, func(a, b int) int {
			for _, v := range p.Rules[a] {
				if v == b {
					return -1
				}
			}
			return 1
		})

		if faulty == reflect.DeepEqual(sorted, update) {
			sum += sorted[(len(sorted)-1)/2]
		}
	}

	return sum
}
