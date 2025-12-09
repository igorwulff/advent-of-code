package shared

import (
	"strings"
)

type Teleporter struct {
	Row   int
	Col   int
	lines []string
}

func (t *Teleporter) Step(row int) (int, bool) {
	if row >= len(t.lines)-1 {
		return 0, true
	}

	split := 0
	cur := t.lines[row]
	next := t.lines[row+1]

	for c, ch := range cur {
		if ch != 'S' && ch != '^' && ch != '|' {
			continue
		}

		if next[c] == '^' {
			split++
			t.lines[row+1] = t.lines[row+1][:c-1] + "|" + string(t.lines[row+1][c]) + "|" + t.lines[row+1][c+2:]
		} else if cur[c] == '|' || cur[c] == 'S' {
			t.lines[row+1] = t.lines[row+1][:c] + "|" + t.lines[row+1][c+1:]
		}
	}

	return split, false
}

func (t Teleporter) Count() int {
	c := 0

	for _, line := range t.lines {
		for _, ch := range line {
			if ch == '|' {
				c++
			}
		}
	}

	return c
}

func ParseInput(input string) Teleporter {
	lines := strings.Split(input, "\n")

	t := Teleporter{
		lines: lines,
	}

	for r, line := range lines {
		for c, ch := range line {
			if ch == 'S' {
				t.Row = r
				t.Col = c
				return t
			}
		}
	}

	return t
}
