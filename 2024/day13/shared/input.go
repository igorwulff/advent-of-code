package shared

import (
	"regexp"
	"strconv"
	"strings"
)

func ParseInput(input string) []Claw {
	lines := strings.Split(input, "\n")

	claws := make([]Claw, 0)

	c := Claw{
		ButtonA: Coord{},
		ButtonB: Coord{},
		Prize:   Coord{},
	}

	regex := regexp.MustCompile(`([A-z\s]*): X[+|=]([0-9]+), Y[+|=]([0-9]+)`)

	for _, line := range lines {
		if line == "" {
			claws = append(claws, c)
			c = Claw{
				ButtonA: Coord{},
				ButtonB: Coord{},
				Prize:   Coord{},
			}

			continue
		}

		matches := regex.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[2])
		y, _ := strconv.Atoi(matches[3])

		switch matches[1] {
		case "Button A":
			c.ButtonA = Coord{X: x, Y: y}
		case "Button B":
			c.ButtonB = Coord{X: x, Y: y}
		case "Prize":
			c.Prize = Coord{X: x, Y: y}
		}
	}

	claws = append(claws, c)
	return claws
}
