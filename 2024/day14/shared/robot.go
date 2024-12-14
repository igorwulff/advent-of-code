package shared

type Robot struct {
	x    int
	y    int
	vx   int
	vy   int
	grid *Grid
}

func (r *Robot) Move() {
	w, h := r.grid.GetDimensions()
	r.x = (r.x + r.vx) % w
	r.y = (r.y + r.vy) % h

	if r.x < 0 {
		r.x += w
	}

	if r.y < 0 {
		r.y += h
	}
}
