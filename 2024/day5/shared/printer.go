package shared

import (
	"reflect"
	"sort"

	"github.com/elliotchance/orderedmap/v3"
)

type Printer struct {
	Rules   map[int][]int
	Updates *orderedmap.OrderedMap[int, []int]
}

func (p Printer) GetSorted(faulty bool) int {
	sum := 0

	for el := p.Updates.Front(); el != nil; el = el.Next() {
		sorted := make([]int, len(el.Value))
		copy(sorted, el.Value)

		sort.Slice(sorted, func(i, j int) bool {
			r := p.Rules[sorted[j]]

			for _, v := range r {
				if v == sorted[i] {
					return false
				}
			}

			return true
		})

		if faulty == reflect.DeepEqual(sorted, el.Value) {
			sum += sorted[(len(sorted)-1)/2]
		}
	}

	return sum
}
