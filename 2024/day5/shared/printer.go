package shared

import "github.com/elliotchance/orderedmap/v3"

type Printer struct {
	Rules   map[int][]int
	Updates *orderedmap.OrderedMap[int, []int]
}
