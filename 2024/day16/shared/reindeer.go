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

type Reindeer struct {
	X    int
	Y    int
	Dir  Dir
	Grid *Grid
}

func (r *Reindeer) NextStep(x, y int, dir Dir) (int, int) {
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

func (r *Reindeer) IsAvailable(x, y int) State {
	if !r.Grid.InBounds(x, y) {
		return Blocked
	}

	cell := r.Grid.GetCell(x, y)
	if cell == "#" {
		return Blocked
	}

	if cell == "E" {
		return Finish
	}

	return Open
}

func (r *Reindeer) Clockwise(dir Dir) Dir {
	return (dir + 1) % 4
}

func (r *Reindeer) CounterClockwise(dir Dir) Dir {
	return (dir + 3) % 4
}

func (r *Reindeer) MoveTowards(start, pos int, dir Dir, tiles *map[int]bool) Dir {
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

func (r *Reindeer) FindPaths(graph *dijkstra.Graph, x, y int, dir Dir) {
	var curPos int = r.Grid.GetPos(x, y)
	var steps uint64 = 1
	for {
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
				graph.AddArc(curPos, leftPos, 1000+steps)
				if leftState == Open {
					r.FindPaths(graph, lx, ly, leftDir)
				}
			}
		}

		if rightState == Open || rightState == Finish {
			if _, ok := vertexes[rightPos]; !ok {
				graph.AddArc(curPos, rightPos, 1000+steps)
				if rightState == Open {
					r.FindPaths(graph, rx, ry, rightDir)
				}
			}
		}

		if nextState == Blocked {
			break
		}

		if nextState == Finish {
			graph.AddArc(curPos, nextPos, steps)
			break
		}

		steps++
		x, y = nx, ny
	}
}
