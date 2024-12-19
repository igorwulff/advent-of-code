package shared

import (
	"sort"
	"strings"
)

type Towel struct {
	patterns []string
	memo     map[string]int
}

func NewTowel(p []string) *Towel {
	// Sort pattern input by length
	sort.Slice(p, func(i, j int) bool {
		return len(p[i]) < len(p[j])
	})

	return &Towel{
		patterns: p,
		memo:     make(map[string]int),
	}
}

func (t *Towel) Match(input string) bool {
	if result, found := t.memo[input]; found {
		return result == 1
	}

	for _, pattern := range t.patterns {
		if pattern == input {
			t.memo[input] = 1
			return true
		}

		if strings.HasSuffix(input, pattern) {
			if t.Match(input[:len(input)-len(pattern)]) {
				t.memo[input] = 1
				return true
			}
		}
	}

	t.memo[input] = 0
	return false
}

func (t *Towel) CountWays(input string) int {
	if count, found := t.memo[input]; found {
		return count
	}

	if input == "" {
		return 1
	}

	i := 0
	for _, pattern := range t.patterns {
		if strings.HasSuffix(input, pattern) {
			i += t.CountWays(input[:len(input)-len(pattern)])
		}
	}

	t.memo[input] = i
	return i
}
