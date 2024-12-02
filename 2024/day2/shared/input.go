package shared

import (
	"strconv"
	"strings"

	"github.com/igorwulff/advent-of-code/utils"
)

type report struct {
	levels []int
}

func (r report) IsSafe(tolerate bool) bool {
	var sortAsc bool

	for i := 0; i < len(r.levels); i++ {
		faulty := false
		levels := r.levels

		if tolerate {
			levels = make([]int, 0)
			levels = append(levels, r.levels[:i]...)
			levels = append(levels, r.levels[i+1:]...)
		}

		// Determine sorting
		if levels[0] < levels[1] {
			sortAsc = true
		} else {
			sortAsc = false
		}

		for i := 1; i < len(levels); i++ {
			if sortAsc {
				if levels[i] < levels[i-1] {
					faulty = true
				}
			} else {
				if levels[i] > levels[i-1] {
					faulty = true
				}
			}

			differ := utils.AbsInt(levels[i] - levels[i-1])
			if differ < 1 || differ > 3 {
				faulty = true
			}
		}

		if !faulty {
			return true
		}

		if !tolerate {
			return false
		}
	}

	return false
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
