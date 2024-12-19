package shared

import (
	"strings"
)

func ParseInput(input string) (*Towel, []string) {
	lines := strings.Split(input, "\n")

	towels := NewTowel(
		strings.Split(lines[0], ", "), // Patterns
	)

	designs := make([]string, 0)

	for k, line := range lines {
		if k == 0 || line == "" {
			continue
		}

		designs = append(designs, line)
	}

	return towels, designs
}
