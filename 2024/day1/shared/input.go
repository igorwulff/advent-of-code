package shared

import (
	"fmt"
	"strings"
)

func ParseInput(input string) ([]int, []int) {
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
