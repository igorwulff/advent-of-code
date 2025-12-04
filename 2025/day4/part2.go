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
		mark := make([]int, 0, 100)

		for y := 0; y < g.Height; y++ {
			for x := 0; x < g.Width; x++ {
				if g.GetCell(x, y) == '.' {
					continue
				}

				if g.CheckAdjacent(x, y, 4) {
					mark = append(mark, g.GetPos(x, y))
					counter++
				}
			}
		}

		if counter == 0 {
			break
		}

		for _, pos := range mark {
			g.Mark(g.GetX(pos), g.GetY(pos))
		}

		sum += counter
	}

	return fmt.Sprint(sum)
}
