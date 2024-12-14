package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day14/shared"
)

var width = 101
var height = 103

// Exported function to be called by the main application
func Part1(input string) string {
	g := shared.NewGrid(width, height)
	robots := shared.ParseInput(input, &g)

	for i := 0; i < 100; i++ {
		for _, robot := range robots {
			robot.Move()
		}
	}

	q := make(map[shared.Quadrant]int)
	for _, r := range robots {
		q[g.GetQuadrant(r)]++
	}

	sum := q[shared.LeftTop] * q[shared.RightTop] * q[shared.LeftBottom] * q[shared.RightBottom]
	return fmt.Sprint(sum)
}
