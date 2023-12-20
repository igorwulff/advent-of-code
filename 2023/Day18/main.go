package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	plan := make(map[string]string, 0)

	minX := 0
	minY := 0
	maxX := 0
	maxY := 0
	x := 0
	y := 0
	plan["X:"+string(x)+",Y:"+string(y)] = "(#000000)"
	for scanner.Scan() {
		raw := scanner.Text()
		text := strings.Split(raw, " ")
		n, _ := strconv.Atoi(text[1])
		hex := text[2]

		if text[0] == "R" {
			i := 0
			for i = x; i < x+n; i++ {
				plan[fmt.Sprintf("X:%d,Y:%d", i, y)] = hex
			}

			x = i
			if x > maxX {
				maxX = x
			}
		} else if text[0] == "L" {
			i := 0
			for i = x; i > x-n; i-- {
				plan[fmt.Sprintf("X:%d,Y:%d", i, y)] = hex
			}

			x = i
			if x < minX {
				minX = x
			}

		} else if text[0] == "D" {
			i := 0
			for i = y; i < y+n; i++ {
				plan[fmt.Sprintf("X:%d,Y:%d", x, i)] = hex
			}

			y = i
			if y > maxY {
				maxY = y
			}
		} else if text[0] == "U" {
			i := 0
			for i = y; i > y-n; i-- {
				plan[fmt.Sprintf("X:%d,Y:%d", x, i)] = hex
			}

			y = i
			if y < minY {
				minY = y
			}
		}
	}

	for y := minY; y <= maxY; y++ {
		fill := false
		for x := minX; x <= maxX; x++ {

			test := fmt.Sprintf("X:%d,Y:%d", x, y)
			if plan[test] == "" {
				if fill {
					fmt.Print("@")
					plan[fmt.Sprintf("X:%d,Y:%d", x, y)] = "1"
					sum++
				} else {
					fmt.Print(".")
				}
			} else {
				value := plan[fmt.Sprintf("X:%d,Y:%d", x, y)]
				if value != "" {
					fill = !fill
					fmt.Print("#")
					sum++
				}
				for xx := x + 1; xx <= maxX; xx++ {
					value := plan[fmt.Sprintf("X:%d,Y:%d", xx, y)]
					if value != "" { // Continous # symbols.
						fmt.Print("#")
						sum++
						x++
					} else {

						valueP := plan[fmt.Sprintf("X:%d,Y:%d", xx+1, y-1)]
						if fill && valueP != "" {
							fill = true
						} else if !fill && valueP == "1" {
							fill = true
						} else {
							fill = false
						}

						break
					}
				}
			}
		}

		fmt.Println()
	}

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

// 137310 Answer is to high...
//That's not the right answer; your answer is too high. If you're stuck, make sure you're using the full input data; there are also some general tips on the about page, or you can ask for hints on the subreddit. Please wait one minute before trying again. [Return to Day 18]
