package shared

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
			r.Perimeter = append(r.Perimeter, g.GetPos(nx, ny))
		} else {
			r.FloodFill(nx, ny)
		}
	}
}

func (r *Region) Sides() int {
	return len(r.Perimeter)
}
