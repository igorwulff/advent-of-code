package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

type node struct {
	value string
	x     int
	y     int
}

var width = 0
var height = 0
var grid []node

func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	height = 0
	for scanner.Scan() {
		for x, row := range scanner.Text() {
			node := node{value: string(row), x: x, y: height}
			grid = append(grid, node)
		}
		if width == 0 {
			width = len(grid)
		}
		height++
	}
	expand()
	stars := findStars()
	sum := 0

	skip := make([]*node, 0)
	for _, star := range stars {
		skip = append(skip, star)
		sum += findClosestNeighbour(star, stars, skip)
	}

	fmt.Printf("Shortest Distances: %d", sum)
	fmt.Println()
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func getPos(x, y int) (*node, error) {
	if x < 0 || x >= width || y < 0 || y >= height {
		return nil, errors.New("Invalid position")
	}
	return &(grid[(y*width)+x]), nil
}

func findClosestNeighbour(star *node, stars []*node, skipStars []*node) int {
	closest := 0
	for _, cell := range stars {
		skip := false
		for _, skipStar := range skipStars {
			if skipStar == cell {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		disX := star.x - cell.x
		disY := star.y - cell.y
		if disX < 0 {
			disX *= -1
		}

		if disY < 0 {
			disY *= -1
		}

		closest += disX + disY
	}

	return closest
}

func findStars() []*node {
	stars := make([]*node, 0)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			cell, _ := getPos(x, y)
			if cell.value == "#" {
				cell.x = x
				cell.y = y
				stars = append(stars, cell)
			}
		}
	}

	return stars
}

func expand() {
	// Check vertical lanes.
	empty := make([]int, 0)
	for x := 0; x < width; x++ {
		// Go from to to bottom.
		isEmpty := true
		for y := 0; y < height; y++ {
			cell, _ := getPos(x, y)
			if cell.value == "#" {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			empty = append(empty, x)
		}
	}

	inc := 1
	for _, x := range empty {
		for y := 0; y < height; y++ {
			grid = slices.Insert(grid, (y*width)+x+inc, node{value: ".", x: x, y: y})
			if y == 1 {
				width++
			}
		}
		inc++
	}

	for y := 0; y < height; y++ {
		// Go from left to right.
		isEmpty := true
		for x := 0; x < width; x++ {
			cell, _ := getPos(x, y)
			if cell.value == "#" {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			for x := 0; x < width; x++ {
				grid = slices.Insert(grid, (y*width)+x, node{value: ".", x: x, y: y})
			}
			height++
			y++
		}
	}
}

func Draw() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell, err := getPos(x, y)
			if err == nil {
				fmt.Print(cell.value + " ")
			} else {
				fmt.Print("? ")
			}
		}

		fmt.Println()
	}
	fmt.Println("--------------------------")
}
