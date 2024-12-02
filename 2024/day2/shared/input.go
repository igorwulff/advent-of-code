package shared

import (
	"strconv"
	"strings"

	"github.com/igorwulff/advent-of-code/utils"
)

type report struct {
	levels []int
}

func (r report) IsSafe() bool {
	var sortAsc bool

	// Determine sorting
	if r.levels[0] < r.levels[1] {
		sortAsc = true
	} else {
		sortAsc = false
	}

	for i := 1; i < len(r.levels); i++ {
		if sortAsc {
			if r.levels[i] < r.levels[i-1] {
				return false
			}
		} else {
			if r.levels[i] > r.levels[i-1] {
				return false
			}
		}

		differ := utils.AbsInt(r.levels[i] - r.levels[i-1])
		if differ < 1 || differ > 3 {
			return false
		}
	}

	return true
}

func ParseInput(input string) []report {
	result := make([]report, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		report := report{
			levels: make([]int, 0),
		}
		for _, v := range strings.Split(line, " ") {
			i, _ := strconv.Atoi(v)
			report.levels = append(report.levels, i)
		}

		result = append(result, report)
	}

	return result
}
