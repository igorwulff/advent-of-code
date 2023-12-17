package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("./sample.txt")
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
		sum += calcLoad(grid, 0, 0, 0, 0)
	}

	// calcLoad(grid, -1, 0, 0, 0, false);
	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func calcLoad(grid *Grid, x, y, nx, ny int) int {
	visited := make([]int, grid.width*grid.height)
	path := make([]int, 0)
	path, visited, _ = movePos(grid, x-1, y, nx, ny, path, visited)

	visited2 := make([]int, grid.width*grid.height)
	path2 := make([]int, 0)
	path2, visited2, _ = movePos(grid, x, y-1, nx, ny, path2, visited2)

	if sum(visited) < sum(visited2) {
		drawPath(grid, visited)
		return sum(visited)
	}
	drawPath(grid, visited2)

	return sum(visited2)
}

func drawPath(grid *Grid, visited []int) {
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			rawPos := getRawPos(grid, x, y)
			if visited[rawPos] > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(grid.cells[rawPos].value)
			}
		}
		fmt.Println()
	}
}

func movePos(grid *Grid, px, py, x, y int, path []int, visited []int) ([]int, []int, error) {
	node, err := getPos(grid, x, y)
	if err != nil {
		return path, visited, err
	}

	rawPos := getRawPos(grid, x, y)
	if visited[rawPos] > 0 {
		return path, visited, errors.New("Already visited")
	}

	maxValue := (x + y) * 6
	sumVisited := sum(visited)
	if sumVisited > maxValue {
		return path, visited, errors.New("To high value.")
	}

	path = append(path, rawPos)
	nodeValue, _ := strconv.Atoi(node.value)
	visited[rawPos] = nodeValue

	// Found end station
	if grid.width-1 == x && grid.height-1 == y {
		return path, visited, nil
	}

	var bestPath []int
	var bestVisited []int

	// Move from left or right
	if px != x {
		cvisited := make([]int, len(visited))
		cpath := make([]int, len(path))
		copy(cvisited, visited)
		copy(cpath, path)

		newPath, visited, err := movePos(grid, x, y, x, y-1, path, visited)
		newPath2, visited2, err2 := movePos(grid, x, y, x, y+1, cpath, cvisited)

		if err == nil || err2 == nil {
			if err != nil && err2 == nil {
				bestPath = newPath2
				bestVisited = visited2
			} else if err2 != nil && err == nil {
				bestPath = newPath
				bestVisited = visited
			} else if sum(visited) < sum(visited2) {
				bestPath = newPath
				bestVisited = visited
			} else {
				bestPath = newPath2
				bestVisited = visited2
			}
		}
	} else if py != y { // Came from top or bottom.
		cvisited := make([]int, len(visited))
		cpath := make([]int, len(path))
		copy(cvisited, visited)
		copy(cpath, path)

		newPath, visited, err := movePos(grid, x, y, x-1, y, path, visited)
		newPath2, visited2, err2 := movePos(grid, x, y, x+1, y, cpath, cvisited)

		if err == nil || err2 == nil {
			if err != nil && err2 == nil {
				bestPath = newPath2
				bestVisited = visited2
			} else if err2 != nil && err == nil {
				bestPath = newPath
				bestVisited = visited
			} else if sum(visited) < sum(visited2) {
				bestPath = newPath
				bestVisited = visited
			} else {
				bestPath = newPath2
				bestVisited = visited2
			}
		}
	}

	// Continue
	isHorizontal := px != x
	canContinue := false
	for i := len(path) - 2; i > len(path)-5; i-- {
		if i < 0 {
			continue
		}

		xx, yy := getPosFromRaw(grid, path[i])
		if isHorizontal && yy != y {
			canContinue = true
			break
		} else if !isHorizontal && xx != x {
			canContinue = true
			break
		}
	}

	if canContinue {
		cvisited := make([]int, len(visited))
		cpath := make([]int, len(path))
		copy(cvisited, visited)
		copy(cpath, path)

		if isHorizontal {
			newPath, visited, err := movePos(grid, x, y, x+x-px, y, cpath, cvisited)

			if err == nil {
				if len(bestPath) > 0 {
					if sum(bestVisited) > sum(visited) {
						bestPath = newPath
						bestVisited = visited
					}
				} else {
					bestPath = newPath
					bestVisited = visited
				}
			}
		} else {
			newPath, visited, err := movePos(grid, x, y, y, y+y-py, cpath, cvisited)

			if err == nil {
				if len(bestPath) > 0 {
					if sum(bestVisited) > sum(visited) {
						bestPath = newPath
						bestVisited = visited
					}
				} else {
					bestPath = newPath
					bestVisited = visited
				}
			}
		}
	}

	if len(bestPath) == 0 {
		return bestPath, visited, errors.New("No entries.")
	}

	return bestPath, bestVisited, nil
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

func sum(list []int) int {
	sum := 0

	for _, v := range list {
		sum += v
	}

	return sum
}

func getPos(grid *Grid, x, y int) (*Node, error) {
	if x < 0 || x >= grid.width || y < 0 || y >= grid.height {
		return nil, errors.New("Invalid position")
	}
	return grid.cells[getRawPos(grid, x, y)], nil
}

func getPosFromRaw(grid *Grid, pos int) (int, int) {
	return pos % grid.width, pos / grid.width
}

func getRawPos(grid *Grid, x, y int) int {
	return (y * grid.width) + x
}
