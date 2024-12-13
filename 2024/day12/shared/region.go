package shared

import (
	"fmt"
)

type Region struct {
	Grid      *Grid
	Value     string
	Positions []int
	Perimeter []int
}

func (r *Region) FloodFill(x, y int) {
	g := r.Grid

	if !g.InBounds(x, y) {
		return
	}

	pos := g.GetPos(x, y)

	if g.Regions[pos] != nil {
		return // Skip already assigned regions
	}

	if g.Cells[pos] != r.Value {
		return // Ignore cells with different values
	}

	r.Positions = append(r.Positions, pos)
	g.Regions[pos] = r

	// Check neighbors to calculate the perimeter
	neighbors := []struct {
		dx, dy int
	}{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // Left, right, top, down
	}

	for _, n := range neighbors {
		nx, ny := x+n.dx, y+n.dy
		if !g.InBounds(nx, ny) || g.Cells[g.GetPos(nx, ny)] != r.Value {
			pos := g.GetPos(nx, ny)
			r.Perimeter = append(r.Perimeter, pos)
		} else {
			r.FloodFill(nx, ny)
		}
	}
}

type Orientation int

const (
	Horizontal Orientation = iota // 0
	Vertical                      // 1
)

func (r *Region) Sides() int {
	sides := 0

	fmt.Println(r.Perimeter)

	for _, pos := range r.Perimeter {
		xx, yy := r.Grid.GetXY(pos)
		fmt.Println("XX", xx, "YY", yy)

		// Check if the position is a corner
		/*corner := false
		for _, n := range r.Perimeter {
			if pos == n {
				continue
			}

			x1, y1 := r.Grid.GetXY(pos)
			x2, y2 := r.Grid.GetXY(n)

			diffx := utils.AbsInt(x1) - utils.AbsInt(x2)
			diffy := utils.AbsInt(y1) - utils.AbsInt(y2)

			if (diffx == 1 || diffx == -1) && (diffy == 1 || diffy == -1) {
				corner = true
				break
			}
		}

		if corner {
			sides++
		}*/
	}

	fmt.Println(sides)

	return sides
}
