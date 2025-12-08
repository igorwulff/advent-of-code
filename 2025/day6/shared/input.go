package shared

import (
	"strconv"
	"strings"
)

type Worksheet struct {
	Problem   []int
	ProblemRL []int
	Modifier  string
}

func (ws Worksheet) CalculateRL() int64 {
	var result int64

	if ws.Modifier == "+" {
		result = 0
		for _, v := range ws.ProblemRL {
			result += int64(v)
		}
	}

	if ws.Modifier == "*" {
		var prod int64 = 1
		for _, v := range ws.ProblemRL {
			prod *= int64(v)
		}
		result = prod
	}

	return result
}

func (ws Worksheet) Calculate() int64 {
	var result int64

	if ws.Modifier == "+" {
		result = 0
		for _, v := range ws.Problem {
			result += int64(v)
		}
	}

	if ws.Modifier == "*" {
		var prod int64 = 1
		for _, v := range ws.Problem {
			prod *= int64(v)
		}
		result = prod
	}

	return result
}

func ParseInput(input string) *[]Worksheet {
	var ws []Worksheet

	lines := strings.Split(input, "\n")
	for l, line := range lines {
		parts := strings.Fields(line)

		for i, part := range parts {
			if l == 0 {
				ws = append(ws, Worksheet{
					Problem:   make([]int, 0, 4),
					ProblemRL: make([]int, 0, 4),
				})
			}

			if part == "*" || part == "+" {
				ws[i].Modifier = part
			} else {
				num, _ := strconv.Atoi(part)
				ws[i].Problem = append(ws[i].Problem, num)
			}
		}
	}

	idx := -1
	last := lines[len(lines)-1]
	for col, val := range last {
		if val == '*' || val == '+' {
			idx++
			ws[idx].Modifier = string(val)
		}

		value := ""
		for i := 0; i < len(lines)-1; i++ {
			char := string(lines[i][col])
			if char == " " {
				continue
			}

			value += char
		}

		if value != "" {
			num, _ := strconv.Atoi(value)
			ws[idx].ProblemRL = append(ws[idx].ProblemRL, num)
		}
	}

	return &ws
}
