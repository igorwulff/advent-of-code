package shared

import (
	"strconv"
	"strings"
)

func ParseInput(input string) ([]string, []int) {
	dir := make([]string, 0)
	steps := make([]int, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		dir = append(dir, string(line[0]))
		var step int
		step, _ = strconv.Atoi(line[1:])
		steps = append(steps, step)
	}

	return dir, steps
}
