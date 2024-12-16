package main

import (
	"fmt"

	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/igorwulff/advent-of-code/2024/day16/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	grid := shared.ParseInput(input)

	graph := dijkstra.NewGraph()
	for i := 0; i < len(grid.Cells); i++ {
		graph.AddEmptyVertex(i)
	}

	grid.Reindeer.FindPaths(&graph, grid.Reindeer.X, grid.Reindeer.Y, shared.East)

	pos1 := grid.GetPos(grid.Reindeer.X, grid.Reindeer.Y)
	pos2 := grid.GetPos(grid.EndTile.X, grid.EndTile.Y)

	path, err := graph.Shortest(pos1, pos2)
	if err != nil {
		panic(err)
	}

	return fmt.Sprint(path.Distance)
}
