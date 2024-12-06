package shared

import "golang.org/x/sync/syncmap"

var (
	Obstacles = syncmap.Map{}
	Width     = 0
	Height    = 0
)

type Grid struct {
	ObstacleX int
	ObstacleY int
}

func (g Grid) InBounds(x, y int) bool {
	return x >= 0 && x < Width && y >= 0 && y < Height
}

func (g *Grid) IsObstacle(x, y int) bool {
	if g.ObstacleX == x && g.ObstacleY == y {
		return true
	}

	if _, ok := Obstacles.Load(g.GetPos(x, y)); ok {
		return true
	}

	return false
}

func (g *Grid) GetX(pos int) int {
	return pos % Width
}

func (g *Grid) GetY(pos int) int {
	return pos / Width
}

func (g *Grid) GetPos(x, y int) int {
	return (Width * y) + x
}
