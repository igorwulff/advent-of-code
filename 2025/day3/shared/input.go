package shared

import (
	"strings"
)

type Ids struct {
	Start int
	End   int
}

func ParseInput(input string) []string {
	output := make([]string, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		output = append(output, line)
	}

	return output
}
