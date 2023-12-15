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
		for cycle := 0; cycle < 4000; cycle++ {
			if cycle%4 == 0 {
				rollNorth(grid)
			} else if cycle%4 == 1 {
				rollWest(grid)
			} else if cycle%4 == 2 {
				rollSouth(grid)
			} else if cycle%4 == 3 {
				rollEast(grid)
			}

		}

		sum += calcLoad(grid)
	}

	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func calcLoad(grid *Grid) int {
	sum := 0
	for y := 0; y < grid.height; y++ {
		row := getRow(grid, y)
		fmt.Println(row)
		for _, x := range row {
			if string(x) == "O" {
				sum += grid.height - y
			}
		}

	}

	return sum
}

func rollNorth(grid *Grid) {
	for x := 0; x < grid.width; x++ {
		for y := 1; y < grid.height; y++ { // Skip first line.
			node, _ := getPos(grid, x, y)
			if node.value == "O" {
				for yy := node.y - 1; yy >= 0; yy-- {
					nextNode, _ := getPos(grid, x, yy)

					if nextNode.value == "." {
						swapPos(grid, node, x, yy)
						node = nextNode
						y--
					} else {
						break
					}
				}
			}
		}
	}
}

func rollSouth(grid *Grid) {
	for x := 0; x < grid.width; x++ {
		for y := grid.height - 1; y >= 0; y-- { // Skip first line.
			node, _ := getPos(grid, x, y)
			if node.value == "O" {
				for yy := node.y + 1; yy < grid.height; yy++ {
					nextNode, _ := getPos(grid, x, yy)

					if nextNode.value == "." {
						swapPos(grid, node, x, yy)
						node = nextNode
					} else {
						break
					}
				}
			}
		}
	}
}

func rollWest(grid *Grid) {
	for y := 0; y < grid.height; y++ {
		for x := 1; x < grid.width; x++ {
			node, _ := getPos(grid, x, y)
			if node.value == "O" {
				for xx := node.x - 1; xx >= 0; x-- {
					nextNode, _ := getPos(grid, xx, y)

					if nextNode.value == "." {
						swapPos(grid, node, xx, y)
						node = nextNode
						x--
					} else {
						break
					}
				}
			}
		}
	}
}

func rollEast(grid *Grid) {
	for y := 0; y < grid.height; y++ {
		for x := grid.width - 1; x >= 0; x-- {
			node, _ := getPos(grid, x, y)
			if node.value == "O" {
				for xx := node.x + 1; xx < grid.width; xx++ {
					nextNode, _ := getPos(grid, xx, y)

					if nextNode.value == "." {
						swapPos(grid, node, xx, y)
						node = nextNode
					} else {
						break
					}
				}
			}
		}
	}
}

func swapPos(grid *Grid, node *Node, x, y int) {
	nextNode, _ := getPos(grid, x, y)
	tmpValue := nextNode.value
	nextNode.value = node.value
	node.value = tmpValue
}

func getPos(grid *Grid, x, y int) (*Node, error) {
	if x < 0 || x >= grid.width || y < 0 || y >= grid.height {
		return nil, errors.New("Invalid position")
	}
	return grid.cells[(y*grid.width)+x], nil
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
