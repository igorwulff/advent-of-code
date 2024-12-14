package shared

import (
	"regexp"
	"strconv"
	"strings"
)

func ParseInput(input string, grid *Grid) []*Robot {
	lines := strings.Split(input, "\n")

	regex := regexp.MustCompile(`([-\d]+)`)

	robots := []*Robot{}

	for _, line := range lines {
		values := regex.FindAllString(line, -1)

		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		vx, _ := strconv.Atoi(values[2])
		vy, _ := strconv.Atoi(values[3])

		robots = append(robots, &Robot{x, y, vx, vy, grid})

	}

	return robots
}
