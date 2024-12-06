package shared

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
		g.SetPos(grid, x, y)
	}

	return true
}

func (g *Guard) SetPos(grid Grid, x, y int) {
	g.X = x
	g.Y = y
	pos := grid.GetPos(x, y)
	g.Visited[pos] = true
	g.Path = append(g.Path, pos)
}
