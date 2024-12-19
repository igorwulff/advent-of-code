package shared

import (
	"sort"
	"strings"
)

type Towel struct {
	patterns []string
}

func NewTowel(p []string) *Towel {
	// Sort pattern input by length
	sort.Slice(p, func(i, j int) bool {
		return len(p[i]) > len(p[j])
	})

	return &Towel{
		patterns: p,
	}
}

func (t *Towel) Match(input string) bool {
	for _, pattern := range t.patterns {
		value := input
		if strings.Contains(value, pattern) {
			value = strings.Replace(value, pattern, "", 1)

			if value == "" {
				return true
			}

			if t.Match(value) {
				return true
			}
		}
	}

	return false
}
