package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2025/day4/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	g := shared.ParseInput(input)

	sum := 0

	for {
		counter := 0
		mark := make([][]bool, g.Height)
		for i := range mark {
			mark[i] = make([]bool, g.Width)
		}

		for y := 0; y < g.Height; y++ {
			for x := 0; x < g.Width; x++ {
				val := g.GetCell(x, y)
				if val == "." {
					continue
				}

				adj := g.CheckAdjacent(x, y)
				if adj < 4 {
					mark[y][x] = true
					counter++
				}
			}
		}

		if counter == 0 {
			break
		}

		for y := 0; y < g.Height; y++ {
			for x := 0; x < g.Width; x++ {
				if mark[y][x] {
					g.Mark(x, y)
				}
			}
		}

		sum += counter
	}

	return fmt.Sprint(sum)
}
