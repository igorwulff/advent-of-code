package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day14/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	g := shared.NewGrid(shared.Width, shared.Height)
	robots := shared.ParseInput(input, &g)

	result := 0
	for i := 1; i < 100000; i++ {
		for _, robot := range robots {
			robot.Move()
		}

		if g.Guess(robots) {
			fmt.Println("#################")
			fmt.Println("Iteration: ", i)
			fmt.Println("#################")
			g.Draw(robots)
			break
		}
	}

	return fmt.Sprint(result)
}
