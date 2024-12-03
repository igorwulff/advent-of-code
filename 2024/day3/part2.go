package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Exported function to be called by the main application
func Part2(input string) string {
	r := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|do\(\)|don't\(\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	output := 0
	enabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
			continue
		}

		if match[0] == "don't()" {
			enabled = false
			continue
		}

		if !enabled {
			continue
		}

		l, _ := strconv.Atoi(match[1])
		r, _ := strconv.Atoi(match[2])
		output += l * r
	}

	return fmt.Sprint(output)
}
