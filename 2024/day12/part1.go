package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day12/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	g := shared.ParseInput(input)

	regions := make([]*shared.Region, 0)
	for pos := range g.Cells {
		if g.Regions[pos] != nil {
			continue
		}

		region := &shared.Region{
			Grid:      &g,
			Value:     g.Cells[pos],
			Positions: make([]int, 0),
			Perimeter: make([]int, 0),
		}
		regions = append(regions, region)

		x, y := g.GetXY(pos)
		region.FloodFill(x, y)
	}

	sum := 0
	for _, region := range regions {
		sum += len(region.Perimeter) * len(region.Positions)
	}

	return fmt.Sprint(sum)
}
