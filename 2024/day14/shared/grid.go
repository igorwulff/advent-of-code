package shared

type Quadrant int

const (
	Middle Quadrant = iota
	LeftTop
	RightTop
	LeftBottom
	RightBottom
)

var Width = 101
var Height = 103

type Grid struct {
	width  int
	height int
}

func NewGrid(width int, height int) Grid {
	return Grid{width, height}
}

func (g *Grid) Draw(robots []*Robot) {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			found := false
			for _, r := range robots {
				if r.x == x && r.y == y {
					found = true
					break
				}
			}

			if found {
				print("#")
			} else {
				print(" ")
			}
		}
		println()
	}
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

func (g *Grid) Guess(robots []*Robot) bool {
	// 50% has a neighbour
	hits := 0.0

	for _, r := range robots {
		for _, r2 := range robots {
			if (r.x == r2.x+1 || r.x == r2.x-1) && (r.y == r2.y+1 || r.y == r2.y-1) {
				hits++
				break
			}
		}
	}

	return hits > float64(len(robots))/2.5
}
