package shared

import "slices"

type Grid struct {
	Width     int
	Height    int
	Guard     Guard
	Obstacles []int
}

func (g Grid) InBounds(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func (g *Grid) IsObstacle(x, y int) bool {
	return slices.Contains(g.Obstacles, g.GetPos(x, y))
}

func (g *Grid) SetObstacle(x, y int) {
	g.Obstacles = append(g.Obstacles, g.GetPos(x, y))
}

func (g *Grid) GetX(pos int) int {
	return pos % g.Width
}

func (g *Grid) GetY(pos int) int {
	return pos / g.Width
}

func (g *Grid) GetPos(x, y int) int {
	return (g.Width * y) + x
}

type Dir int

const (
	Top Dir = iota
	Right
	Bottom
	Left
)

type Guard struct {
	X      int
	Y      int
	CurDir Dir
	Path   []int
}

func (g *Guard) Move(grid Grid) bool {
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
		return false
	}

	if grid.IsObstacle(x, y) {
		g.CurDir = (g.CurDir + 1) % 4
	} else {
		g.X = x
		g.Y = y
		g.Path = append(g.Path, grid.GetPos(x, y))
	}

	return true
}

func (g *Guard) GetVisited() []int {
	visited := make([]int, 0)

	for _, pos := range g.Path {
		if !slices.Contains(visited, pos) {
			visited = append(visited, pos)
		}
	}

	return visited
}
