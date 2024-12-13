package shared

import "fmt"

type Coord struct {
	X int
	Y int
}

type Claw struct {
	ButtonA Coord
	ButtonB Coord
	Prize   Coord
}

type Cost int

func (c Claw) Matches() []int {
	b1X, b1Y := c.ButtonA.X, c.ButtonA.Y
	b2X, b2Y := c.ButtonB.X, c.ButtonB.Y
	prizeX, prizeY := c.Prize.X, c.Prize.Y

	maxA := prizeX / b1X
	matches := []int{}

	for buttonA := maxA; buttonA > 0; buttonA-- {
		remainingX := prizeX - b1X*buttonA
		if remainingX%b2X != 0 {
			continue
		}

		buttonB := remainingX / b2X
		if buttonB*b2Y == prizeY-b1Y*buttonA {
			fmt.Print("Match: ", buttonA, buttonB, "\n")

			matches = append(matches, buttonA*3+buttonB)
		}
	}

	return matches
}
