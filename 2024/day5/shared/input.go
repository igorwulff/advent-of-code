package shared

import (
	"strconv"
	"strings"
)

func ParseInput(input string) Printer {
	printer := Printer{
		Rules:   make(map[int][]int),
		Updates: make([][]int, 0),
	}

	lines := strings.Split(input, "\n")
	ordering := false
	updates := 0
	for _, line := range lines {
		if line == "" {
			ordering = true
			continue
		}

		if !ordering {
			parts := strings.Split(line, "|")
			l, _ := strconv.Atoi(parts[0])
			r, _ := strconv.Atoi(parts[1])

			printer.Rules[l] = append(printer.Rules[l], r)
		} else {
			partsValues := strings.Split(line, ",")

			parts := make([]int, len(partsValues))
			for i, part := range partsValues {
				parts[i], _ = strconv.Atoi(part)
			}

			printer.Updates = append(printer.Updates, parts)
			updates++
		}
	}

	return printer
}
