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
	Visited map[int]bool
}

func (g *Guard) Reset(grid Grid, x, y int) {
	g.CurDir = Top
	g.Path = make([]int, 0)
	g.Visited = make(map[int]bool)
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
	} else {
		g.SetPos(grid, x, y)
	}

	if len(g.Path) > 10000 {
		return false, fmt.Errorf("Guard got stuck")
	}

	return true, nil
}

func (g *Guard) SetPos(grid Grid, x, y int) {
	g.X = x
	g.Y = y
	pos := grid.GetPos(x, y)
	g.Visited[pos] = true
	g.Path = append(g.Path, pos)
}
