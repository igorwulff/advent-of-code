package shared

import (
	"github.com/igorwulff/advent-of-code/utils"
)

type Dir int

const (
	North Dir = iota
	East
	South
	West
)

type State int

const (
	Blocked State = iota
	Open
	Finish
)

type User struct {
	X    int
	Y    int
	Grid *Grid
}

func (r *User) NextStep(x, y int, dir Dir) (int, int) {
	switch dir {
	case North:
		y--
	case East:
		x++
	case South:
		y++
	case West:
		x--
	}

	return x, y
}

func (r *User) IsAvailable(x, y int) State {
	if !r.Grid.InBounds(x, y) {
		return Blocked
	}

	if x == r.Grid.EndTile.X && y == r.Grid.EndTile.Y {
		return Finish
	}

	cell := r.Grid.GetCell(x, y)
	if cell == "#" {
		return Blocked
	}

	return Open
}

func (r *User) Clockwise(dir Dir) Dir {
	return (dir + 1) % 4
}

func (r *User) CounterClockwise(dir Dir) Dir {
	return (dir + 3) % 4
}

func (r *User) MoveTowards(start, pos int, dir Dir, tiles *map[int]bool) Dir {
	(*tiles)[start] = true
	(*tiles)[pos] = true

	x1, y1 := r.Grid.GetXY(start)
	x2, y2 := r.Grid.GetXY(pos)

	for {
		for _, dir := range []Dir{North, East, South, West} {
			nx, ny := r.NextStep(x1, y1, dir)
			if nx == x2 && ny == y2 {
				(*tiles)[r.Grid.GetPos(nx, ny)] = true
				return dir
			}
		}

		x1, y1 = r.NextStep(x1, y1, dir)
		(*tiles)[r.Grid.GetPos(x1, y1)] = true
	}
}

func (r *User) FindPath(x, y int, dir Dir) []int {
	tiles := make([]int, 0)

	tiles = append(tiles, r.Grid.GetPos(x, y))

	for r.IsAvailable(x, y) != Finish {
		if dir != East && r.IsAvailable(x-1, y) != Blocked {
			dir = West
		} else if dir != West && r.IsAvailable(x+1, y) != Blocked {
			dir = East
		} else if dir != South && r.IsAvailable(x, y-1) != Blocked {
			dir = North
		} else if dir != North && r.IsAvailable(x, y+1) != Blocked {
			dir = South
		}

		x, y = r.NextStep(x, y, dir)

		tiles = append(tiles, r.Grid.GetPos(x, y))
	}

	return tiles
}

func (r *User) FindCheats(path []int, cheats int) int {
	sum := 0
	max := len(path)
	for idx, pos := range path {
		x, y := r.Grid.GetXY(pos)

		min := idx + PicoSeconds + 1
		if min >= max {
			break
		}

		for i := min; i < max; i++ {
			x2, y2 := r.Grid.GetXY(path[i])
			dist := utils.AbsInt(x-x2) + utils.AbsInt(y-y2)

			totalDistance := idx + dist + max - i

			if dist <= cheats && (max-totalDistance) >= PicoSeconds {
				sum++
			}
		}
	}

	return sum
}
