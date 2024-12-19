package shared

import (
	"sort"
	"strings"
)

type Towel struct {
	patterns  []string
	memo      map[string]bool
	memoCount map[string]int
}

func NewTowel(p []string) *Towel {
	// Sort pattern input by length
	sort.Slice(p, func(i, j int) bool {
		return len(p[i]) > len(p[j])
	})

	return &Towel{
		patterns:  p,
		memo:      make(map[string]bool),
		memoCount: make(map[string]int),
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

func (t *Towel) CountWays(input string) int {
	// Check if the result is already cached
	if count, found := t.memoCount[input]; found {
		return count
	}

	// If the input is empty, there is exactly one way (no patterns used)
	if input == "" {
		return 1
	}

	ways := 0

	// Check all patterns
	for _, pattern := range t.patterns {
		// If the pattern matches the end of the input
		if strings.HasSuffix(input, pattern) {
			// Count ways for the remaining input
			ways += t.CountWays(input[:len(input)-len(pattern)])
		}
	}

	// Cache and return the result
	t.memoCount[input] = ways
	return ways
}
