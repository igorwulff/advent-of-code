package shared

type Grid struct {
	Width     int
	Height    int
	Obstacles map[int]bool
}

func (g Grid) InBounds(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func (g *Grid) IsObstacle(x, y int) bool {
	return g.Obstacles[g.GetPos(x, y)]
}

func (g *Grid) SetObstacle(x, y int) {
	g.Obstacles[g.GetPos(x, y)] = true
}

func (g *Grid) RemoveObstacle(x, y int) {
	delete(g.Obstacles, g.GetPos(x, y))
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
