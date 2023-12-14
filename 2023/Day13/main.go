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
		if !grid.isVertical {
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
	getReflectionLine(grid, true)
	getReflectionLine(grid, false)

	for xx := 0; xx < grid.width; xx++ {
		for yy := 0; yy < grid.height; yy++ {
			node, _ := getPos(grid, xx, yy)
			tmpValue := node.value
			tmpReflectionLine := grid.reflectionLine
			tmpIsVertical := grid.isVertical

			if tmpValue == "#" {
				node.value = "."
			} else {
				node.value = "#"
			}

			getReflectionLine(grid, true)
			if grid.reflectionLine != tmpReflectionLine || grid.isVertical != tmpIsVertical {
				return
			}

			getReflectionLine(grid, false)
			if grid.reflectionLine != tmpReflectionLine || grid.isVertical != tmpIsVertical {
				return
			}

			// Reset values for next iteration.
			grid.reflectionLine = tmpReflectionLine
			grid.isVertical = tmpIsVertical
			node.value = tmpValue
		}
	}
}

func getReflectionLine(grid *Grid, isVertical bool) {
	var size int

	if isVertical {
		size = grid.height
	} else {
		size = grid.width
	}

	for i := 1; i < size; i++ {
		ii := 1
		for {
			var nodeLeft, nodeRight string

			if i-ii >= 0 && i+ii <= size {
				if isVertical {
					nodeLeft = getRow(grid, i-ii)
					nodeRight = getRow(grid, i+ii-1)
				} else {
					nodeLeft = getColumn(grid, i-ii)
					nodeRight = getColumn(grid, i+ii-1)
				}
			}

			if i-ii < 0 || i+ii > size {
				if grid.reflectionLine == 0 || grid.reflectionLine != i || grid.isVertical != isVertical {
					grid.reflectionLine = i
					grid.isVertical = isVertical
					return
				} else {
					break
				}
			} else if i-ii >= 0 && i+ii <= size && nodeLeft == nodeRight {
				ii++
			} else {
				break
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

/*
45894
That's not the right answer; your answer is too high
. If you're stuck, make sure you're using the full input data; there are also some general tips on the about page, or you can ask for hints on the subreddit.
 Please wait one minute before trying again. [Return to Day 13]
*/
