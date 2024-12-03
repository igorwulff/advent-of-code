package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Exported function to be called by the main application
func Part1(input string) string {
	r, err := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)

	if err != nil {
		return "Regex can not be compiled."
	}

	matches := r.FindAllStringSubmatch(input, -1)

	output := 0
	for _, match := range matches {
		l, _ := strconv.Atoi(match[1])
		r, _ := strconv.Atoi(match[2])
		output += l * r
	}

	return fmt.Sprint(output)
}
