package main

import (
	"fmt"

	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/igorwulff/advent-of-code/2024/day18/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	if shared.CorruptedBytes == 0 {
		shared.CorruptedBytes = 1024
	}

	if shared.Width == 0 || shared.Height == 0 {
		shared.Width = 71
		shared.Height = 71
	}

	grid := shared.ParseInput(input)

	graph := dijkstra.NewGraph()
	for i := 0; i < len(grid.Cells); i++ {
		graph.AddEmptyVertex(i)
	}

	grid.User.FindPaths(&graph, grid.User.X, grid.User.Y, shared.East)
	pos1 := grid.GetPos(grid.User.X, grid.User.Y)
	pos2 := grid.GetPos(grid.EndTile.X, grid.EndTile.Y)

	path, err := graph.Shortest(pos1, pos2)
	if err != nil {
		panic(err)
	}

	return fmt.Sprint(path.Distance)
}
