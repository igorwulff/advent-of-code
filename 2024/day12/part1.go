package main

import (
	"fmt"
	"strings"
)

type Grid struct {
	Width   int
	Height  int
	Cells   []string
	Regions []*Region
}

func (g Grid) InBounds(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func (g *Grid) GetX(pos int) int {
	return pos % g.Width
}

func (g *Grid) GetY(pos int) int {
	if pos < 0 {
		abs := -pos
		return (abs / g.Width) * -1
	}
	return pos / g.Width
}

func (g Grid) GetXY(pos int) (int, int) {
	return g.GetX(pos), g.GetY(pos)
}

func (g *Grid) GetPos(x, y int) int {
	return (g.Width * y) + x
}

type Region struct {
	Value     string
	Positions []int
	Perimeter int
}

// Exported function to be called by the main application
func Part1(input string) string {
	lines := strings.Split(input, "\n")

	g := Grid{
		Width:   len(lines[0]),
		Height:  len(lines),
		Cells:   make([]string, 0, len(lines[0])*len(lines)),
		Regions: make([]*Region, len(lines[0])*len(lines)),
	}

	for _, line := range lines {
		g.Cells = append(g.Cells, strings.Split(line, "")...)
	}

	regions := make([]*Region, 0)
	for pos := range g.Cells {
		if g.Regions[pos] != nil {
			continue
		}

		region := &Region{
			Value:     g.Cells[pos],
			Positions: make([]int, 0),
			Perimeter: 0,
		}
		regions = append(regions, region)

		x, y := g.GetXY(pos)
		floodFill(&g, region, x, y)
	}

	sum := 0
	for _, region := range regions {
		sum += region.Perimeter * len(region.Positions)
	}

	return fmt.Sprint(sum)
}

func floodFill(g *Grid, region *Region, x, y int) {
	if !g.InBounds(x, y) {
		return
	}

	pos := g.GetPos(x, y)

	if g.Regions[pos] != nil {
		return // Skip already assigned regions
	}

	if g.Cells[pos] != region.Value {
		return // Ignore cells with different values
	}

	region.Positions = append(region.Positions, pos)
	g.Regions[pos] = region

	// Check neighbors to calculate the perimeter
	neighbors := []struct {
		dx, dy int
	}{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // Left, right, top, down
	}

	for _, n := range neighbors {
		nx, ny := x+n.dx, y+n.dy
		if !g.InBounds(nx, ny) || g.Cells[g.GetPos(nx, ny)] != region.Value {
			region.Perimeter++
		} else {
			floodFill(g, region, nx, ny)
		}
	}
}
