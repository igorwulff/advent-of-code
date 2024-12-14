package shared

type Quadrant int

const (
	Middle Quadrant = iota
	LeftTop
	RightTop
	LeftBottom
	RightBottom
)

type Grid struct {
	width  int
	height int
}

func NewGrid(width int, height int) Grid {
	return Grid{width, height}
}

func (g *Grid) GetQuadrant(robot *Robot) Quadrant {
	x := robot.x
	y := robot.y

	w := (g.width - 1) / 2
	h := (g.height - 1) / 2

	if x < w && y < h {
		return LeftTop
	} else if x > w && y < h {
		return RightTop
	} else if x < w && y > h {
		return LeftBottom
	} else if x > w && y > h {
		return RightBottom
	}

	return Middle
}

func (g *Grid) GetDimensions() (int, int) {
	return g.width, g.height
}
