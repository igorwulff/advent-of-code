package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	grids := make([]*Grid, 0)

	nodes := make([]*Node, 0)
	grid := Grid{width: 0, height: 0, cells: nodes}
	grids = append(grids, &grid)

	text := ""
	for scanner.Scan() {
		text = scanner.Text()
		grid := grids[len(grids)-1]
		if text == "" { // Next Line
			nodes := make([]*Node, 0)
			grids = append(grids, &Grid{width: 0, height: 0, cells: nodes})
			continue
		}

		for x, row := range text {
			grid.cells = append(grid.cells, &Node{value: string(row), x: x, y: grid.height})
		}
		if grid.width == 0 {
			grid.width = len(grid.cells)
		}
		grid.height++
	}

	for _, grid := range grids {
		for xx := 0; xx < grid.width; xx++ {
			tmpSum := calcLoad(grid, xx, -1, xx, 0)
			if tmpSum > sum {
				sum = tmpSum
			}
			tmpSum = calcLoad(grid, xx, grid.height, xx, grid.height-1)
			if tmpSum > sum {
				sum = tmpSum
			}
		}

		for yy := 0; yy < grid.height; yy++ {
			tmpSum := calcLoad(grid, -1, yy, 0, yy)
			if tmpSum > sum {
				sum = tmpSum
			}
			tmpSum = calcLoad(grid, grid.width, yy, grid.width-1, yy)
			if tmpSum > sum {
				sum = tmpSum
			}
		}
	}

	// calcLoad(grid, -1, 0, 0, 0, false);
	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func resetTouched(grid *Grid) {
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			node, _ := getPos(grid, x, y)
			node.touched = false
		}
	}
}

func calcLoad(grid *Grid, x, y, nx, ny int) int {
	moveLight(grid, x, y, nx, ny, 0)

	sum := 0
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			node, _ := getPos(grid, x, y)
			if node.touched {
				fmt.Print("#")
				sum++
			} else {
				fmt.Print(node.value)
			}
		}
		fmt.Println()
	}

	resetTouched(grid)

	return sum
}

func moveLight(grid *Grid, x, y, nx, ny, repeat int) {
	node, err := getPos(grid, nx, ny)
	if err != nil || repeat > 100 {
		return
	}

	if node.touched {
		repeat++
	} else {
		repeat = 0
	}

	node.touched = true

	passthrough := false
	if node.value == "|" {
		if nx-x != 0 { // Comes from left or right.
			moveLight(grid, nx, ny, nx, ny-1, repeat)
			moveLight(grid, nx, ny, nx, ny+1, repeat)
		} else {
			passthrough = true
		}
	} else if node.value == "-" {
		if ny-y != 0 { // Comes from top or bottom
			moveLight(grid, nx, ny, nx-1, ny, repeat)
			moveLight(grid, nx, ny, nx+1, ny, repeat)
		} else {
			passthrough = true
		}
	} else if node.value == "/" {
		if nx < x { // Comes from right
			moveLight(grid, nx, ny, nx, ny+1, repeat)
		} else if nx > x { // Comes from left
			moveLight(grid, nx, ny, nx, ny-1, repeat)
		} else if ny < y { // Comes from bottom
			moveLight(grid, nx, ny, nx+1, ny, repeat)
		} else if ny > y { // Comes from top
			moveLight(grid, nx, ny, nx-1, ny, repeat)
		}
	} else if node.value == "\\" {
		if nx < x { // Comes from right
			moveLight(grid, nx, ny, nx, ny-1, repeat)
		} else if nx > x { // Comes from left
			moveLight(grid, nx, ny, nx, ny+1, repeat)
		} else if ny < y { // Comes from bottom
			moveLight(grid, nx, ny, nx-1, ny, repeat)
		} else if ny > y { // Comes from top
			moveLight(grid, nx, ny, nx+1, ny, repeat)
		}
	}

	if node.value == "." || passthrough {
		if nx > x {
			moveLight(grid, nx, ny, nx+1, ny, repeat)
		} else if nx < x {
			moveLight(grid, nx, ny, nx-1, ny, repeat)
		} else if ny > y {
			moveLight(grid, nx, ny, nx, ny+1, repeat)
		} else if ny < y {
			moveLight(grid, nx, ny, nx, ny-1, repeat)
		}
	}
}

type Grid struct {
	width  int
	height int
	cells  []*Node
}

type Node struct {
	value   string
	touched bool
	x       int
	y       int
}

func getColumn(grid *Grid, x int) string {
	row := ""

	for y := 0; y < grid.height; y++ {
		node, _ := getPos(grid, x, y)
		row += node.value
	}
	return row
}

func getRow(grid *Grid, y int) string {
	row := ""

	for x := 0; x < grid.width; x++ {
		node, _ := getPos(grid, x, y)
		row += node.value
	}

	return row
}

func getPos(grid *Grid, x, y int) (*Node, error) {
	if x < 0 || x >= grid.width || y < 0 || y >= grid.height {
		return nil, errors.New("Invalid position")
	}
	return grid.cells[(y*grid.width)+x], nil
}
