package shared

import "fmt"

type Dir int

const (
	Top Dir = iota
	Right
	Bottom
	Left
)

type Guard struct {
	X       int
	Y       int
	CurDir  Dir
	Path    []int
	Visited map[int]Dir
}

func (g *Guard) Reset(grid Grid, x, y int) {
	g.CurDir = Top
	g.Path = make([]int, 0)
	g.Visited = make(map[int]Dir)
	g.SetPos(grid, x, y)
}

func (g *Guard) Move(grid Grid) (bool, error) {
	x := g.X
	y := g.Y

	switch g.CurDir {
	case Top:
		y--
	case Right:
		x++
	case Bottom:
		y++
	case Left:
		x--
	}

	if !grid.InBounds(x, y) {
		return false, nil
	}

	if grid.IsObstacle(x, y) {
		g.CurDir = (g.CurDir + 1) % 4
		return true, nil
	}

	if val, ok := g.Visited[grid.GetPos(x, y)]; ok {
		if val == g.CurDir {
			return false, fmt.Errorf("guard got stuck")
		}
	}

	g.SetPos(grid, x, y)

	return true, nil
}

func (g *Guard) SetPos(grid Grid, x, y int) {
	g.X = x
	g.Y = y
	pos := grid.GetPos(x, y)
	g.Visited[pos] = g.CurDir
	g.Path = append(g.Path, pos)
}
