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

	pos := grid.GetPos(x, y)

	if val, ok := g.Visited[pos]; ok {
		if val == g.CurDir {
			return false, fmt.Errorf("guard got stuck")
		}
	}

	g.SetPos(x, y, pos)

	return true, nil
}

func (g *Guard) SetPos(x, y, pos int) {
	g.X = x
	g.Y = y
	g.Visited[pos] = g.CurDir
	g.Path = append(g.Path, pos)
}
