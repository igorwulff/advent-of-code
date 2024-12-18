package shared

import (
	"github.com/RyanCarrier/dijkstra/v2"
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

func (r *User) FindPaths(graph *dijkstra.Graph, x, y int, dir Dir) {
	curPos := r.Grid.GetPos(x, y)
	vertexes, _ := graph.GetVertexArcs(curPos)

	nx, ny := r.NextStep(x, y, dir)
	nextState := r.IsAvailable(nx, ny)
	nextPos := r.Grid.GetPos(nx, ny)

	leftDir := r.CounterClockwise(dir)
	lx, ly := r.NextStep(x, y, leftDir)
	leftState := r.IsAvailable(lx, ly)
	leftPos := r.Grid.GetPos(lx, ly)

	rightDir := r.Clockwise(dir)
	rx, ry := r.NextStep(x, y, rightDir)
	rightState := r.IsAvailable(rx, ry)
	rightPos := r.Grid.GetPos(rx, ry)

	if leftState == Open || leftState == Finish {
		if _, ok := vertexes[leftPos]; !ok {
			graph.AddArc(curPos, leftPos, 1)
			if leftState == Open {
				r.FindPaths(graph, lx, ly, leftDir)
			}
		}
	}

	if rightState == Open || rightState == Finish {
		if _, ok := vertexes[rightPos]; !ok {
			graph.AddArc(curPos, rightPos, 1)
			if rightState == Open {
				r.FindPaths(graph, rx, ry, rightDir)
			}
		}
	}

	if nextState == Open || nextState == Finish {
		if _, ok := vertexes[nextPos]; !ok {
			graph.AddArc(curPos, nextPos, 1)
			if nextState == Open {
				r.FindPaths(graph, nx, ny, dir)
			}
		}
	}
}
