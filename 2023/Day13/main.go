package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type Grid struct {
	width          int
	height         int
	cells          []*Node
	reflectionLine int
	isVertical     bool
}

type Node struct {
	value string
	x     int
	y     int
}

func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grids := make([]*Grid, 0)

	nodes := make([]*Node, 0)
	grid := Grid{width: 0, height: 0, cells: nodes, isVertical: true}
	grids = append(grids, &grid)

	text := ""
	for scanner.Scan() {
		text = scanner.Text()
		grid := grids[len(grids)-1]
		if text == "" {
			nodes := make([]*Node, 0)
			grids = append(grids, &Grid{width: 0, height: 0, cells: nodes, isVertical: true})
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

	sum := 0
	for _, grid := range grids {
		findReflection(grid)

		if grid.isVertical {
			sum += grid.reflectionLine
		} else {
			sum += grid.reflectionLine * 100
		}
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func getPos(grid *Grid, x, y int) (*Node, error) {
	if x < 0 || x >= grid.width || y < 0 || y >= grid.height {
		return nil, errors.New("Invalid position")
	}
	return grid.cells[(y*grid.width)+x], nil
}

func findReflection(grid *Grid) {
outX:
	for x := 1; x < grid.width; x++ {
		i := 1
		for {
			if x-i < 0 || x+i > grid.width { // Improve i > 2 to check how many should be the same.
				grid.reflectionLine = x
				grid.isVertical = true
				continue outX
			} else if x-i >= 0 && x+i <= grid.width && getColumn(grid, x-i) == getColumn(grid, x+i-1) {
				i++
			} else {
				break
			}
		}
	}

	if grid.reflectionLine == 0 {
	outY:
		for y := 1; y < grid.height; y++ {
			i := 1

			for {
				if y-i < 0 || y+i > grid.height {
					// is reflection
					grid.reflectionLine = y
					grid.isVertical = false
					continue outY
				} else if y-i >= 0 && y+i <= grid.height && getRow(grid, y-i) == getRow(grid, y+i-1) {
					i++
				} else {
					break
				}
			}
		}
	}
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
