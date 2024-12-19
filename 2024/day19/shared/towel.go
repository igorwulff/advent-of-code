package shared

import (
	"sort"
	"strings"
)

type Towel struct {
	patterns []string
	memo     map[string]bool
}

func NewTowel(p []string) *Towel {
	// Sort pattern input by length
	sort.Slice(p, func(i, j int) bool {
		return len(p[i]) > len(p[j])
	})

	return &Towel{
		patterns: p,
		memo:     make(map[string]bool),
	}
}

func (t *Towel) Match(input string) bool {
	// Check if the result is already cached
	if result, found := t.memo[input]; found {
		return result
	}

	for _, pattern := range t.patterns {
		if pattern == input {
			// Cache and return the result
			t.memo[input] = true
			return true
		}

		// Check if the pattern is at the end of the input
		if strings.HasSuffix(input, pattern) {
			// Check if the rest of the input is valid

			if t.Match(input[:len(input)-len(pattern)]) {
				// Cache and return the result
				t.memo[input] = true
				return true
			}
		}
	}

	// Cache and return the result
	t.memo[input] = false
	return false
}
