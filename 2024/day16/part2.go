package main

import (
	"fmt"

	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/igorwulff/advent-of-code/2024/day16/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	grid := shared.ParseInput(input)

	graph := dijkstra.NewGraph()
	for i := 0; i < len(grid.Cells); i++ {
		graph.AddEmptyVertex(i)
	}

	grid.Reindeer.FindPaths(&graph, grid.Reindeer.X, grid.Reindeer.Y, shared.East)

	pos1 := grid.GetPos(grid.Reindeer.X, grid.Reindeer.Y)
	pos2 := grid.GetPos(grid.EndTile.X, grid.EndTile.Y)

	paths, err := graph.ShortestAll(pos1, pos2)
	if err != nil {
		panic(err)
	}

	tiles := make(map[int]bool)
	for _, path := range paths.Paths {
		dir := shared.East
		start := grid.GetPos(grid.Reindeer.X, grid.Reindeer.Y)

		for k, step := range path {
			if k == 0 {
				continue
			}
			dir = grid.Reindeer.MoveTowards(start, step, dir, &tiles)
			start = step
		}

	}

	return fmt.Sprint(len(tiles))
}
