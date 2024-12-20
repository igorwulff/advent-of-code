package main

import (
	"fmt"

	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/igorwulff/advent-of-code/2024/day20/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
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

	fastestTime := path.Distance

	sum := 0
	prevPos := -1
	for _, pos := range grid.Walls {
		x, y := grid.GetXY(pos)

		if x == 0 || y == 0 || x == grid.Width-1 || y == grid.Height-1 {
			continue
		}

		grid.AssignCheatWall(pos, prevPos)
		prevPos = pos

		graph := dijkstra.NewGraph()
		for i := 0; i < len(grid.Cells); i++ {
			graph.AddEmptyVertex(i)
		}

		grid.User.FindPaths(&graph, grid.User.X, grid.User.Y, shared.East)
		path, err := graph.Shortest(pos1, pos2)
		if err != nil {
			continue
		}

		if path.Distance < fastestTime && fastestTime-path.Distance >= uint64(shared.PicoSeconds) {
			sum++
		}
	}

	return fmt.Sprint(sum)
}
