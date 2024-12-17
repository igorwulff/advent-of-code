package shared

import (
	"regexp"
	"strconv"
	"strings"
)

func ParseInput(input string) *Computer {
	lines := strings.Split(input, "\n")

	comp := Computer{
		Registers: map[string]int{
			"A": 0,
			"B": 0,
			"C": 0,
		},
		Pointer: 0,
	}

	regex := regexp.MustCompile(`([A-z]*): ([0-9,]*)`)
	for _, line := range lines {
		if line == "" {
			continue
		}

		matches := regex.FindStringSubmatch(line)

		switch matches[1] {
		case "A", "B", "C":
			value, _ := strconv.Atoi(matches[2])
			comp.Registers[matches[1]] = value

		case "Program":
			comp.Program = make([]int, 0)
			values := strings.Split(matches[2], ",")
			for _, v := range values {
				value, _ := strconv.Atoi(v)
				comp.Program = append(comp.Program, value)
			}
		}
	}

	return &comp
}
