package shared

import (
	"strconv"
	"strings"
)

type Worksheet struct {
	Problem  []int
	Modifier string
}

func (ws Worksheet) Calculate() int64 {
	var result int64

	switch ws.Modifier {
	case "+":
		result = 0
		for _, v := range ws.Problem {
			result += int64(v)
		}
	case "-":
		result = int64(ws.Problem[0])
		for i := 1; i < len(ws.Problem); i++ {
			result -= int64(ws.Problem[i])
		}
	case "*":
		var prod int64 = 1
		for _, v := range ws.Problem {
			prod *= int64(v)
		}
		result = prod
	case "/":
		result = int64(ws.Problem[0])
		for i := 1; i < len(ws.Problem); i++ {
			result /= int64(ws.Problem[i])
		}
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
					Problem: make([]int, 0, 4),
				})
			}

			if part == "*" || part == "+" || part == "-" || part == "/" {
				ws[i].Modifier = part
			} else {
				num, _ := strconv.Atoi(part)
				ws[i].Problem = append(ws[i].Problem, num)
			}
		}
	}

	return &ws
}
