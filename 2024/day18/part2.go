package main

import (
	"fmt"

	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/igorwulff/advent-of-code/2024/day18/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	if shared.CorruptedBytes == 0 {
		shared.CorruptedBytes = 1024
	}

	if shared.Width == 0 || shared.Height == 0 {
		shared.Width = 71
		shared.Height = 71
	}

	grid := shared.ParseInput(input)

	var x, y int
	for {
		pos, err := grid.PlaceCorruptedByte()
		if err != nil {
			break
		}
		x, y = grid.GetXY(pos)

		graph := dijkstra.NewGraph()
		for i := 0; i < len(grid.Cells); i++ {
			graph.AddEmptyVertex(i)
		}

		grid.User.FindPaths(&graph, grid.User.X, grid.User.Y, shared.East)
		pos1 := grid.GetPos(grid.User.X, grid.User.Y)
		pos2 := grid.GetPos(grid.EndTile.X, grid.EndTile.Y)

		_, err = graph.Shortest(pos1, pos2)
		if err != nil {
			break
		}
	}

	return fmt.Sprintf("%d,%d", x, y)
}
