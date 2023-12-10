package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

var strength = map[string]int{"A": 14, "K": 13, "Q": 12, "J": 1, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2}

type node struct {
	value string
	x     int
	y     int
	prev  *node
	next  *node
}

var width = 0
var height = 0
var grid []node

func main() {
	start := time.Now()

	file, err := os.Open("./sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var startNode node
	y := 0
	for scanner.Scan() {
		for x, row := range scanner.Text() {
			node := node{value: string(row), x: x, y: y}
			if row == 'S' {
				startNode = node
			}
			grid = append(grid, node)
		}
		if width == 0 {
			width = len(grid)
		}
		y++
	}
	height = y
	longest := 0
	starts := findStartingPositions(startNode)
	for _, start := range starts {
		next := start
		steps := 0
		for {
			next, err = findNextPosition(next)
			if err != nil || string(next.value) == "S" {
				break
			} else {
				steps++
			}
		}
		if string(next.value) == "S" && steps > longest {
			longest = steps
		}
	}

	fmt.Println((longest / 2) + 1)
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func getPos(x, y int) (*node, error) {
	if x < 0 || x > width || y < 0 || y > height {
		return nil, errors.New("Invalid position")
	}
	return &(grid[(y*width)+x]), nil
}

func findStartingPositions(start node) []node {
	nodes := make([]node, 0)
	for x := start.x - 1; x <= start.x+1; x++ {
		for y := start.y - 1; y <= start.y+1; y++ {
			if x == start.x && y == start.y {
				continue
			}

			curNode, _ := getPos(x, y)
			if curNode == nil || curNode.value == "." {
				continue
			}

			curNode.prev = &start

			nodes = append(nodes, *curNode)
		}
	}

	return nodes
}

func findNextPosition(cur node) (node, error) {
	var nextNode *node

	if cur.value == "|" {
		if cur.x == cur.prev.x { // Same horizontal axis.
			nextY := cur.y + 1
			if cur.y < cur.prev.y {
				nextY = cur.y - 1
			}
			nextNode, _ = getPos(cur.x, nextY) // 0+(0-1) = -1*-1
		}
	} else if cur.value == "-" {
		if cur.y == cur.prev.y { // Same vertical axis.
			nextX := cur.x + 1
			if cur.x < cur.prev.x {
				nextX = cur.x - 1
			}
			nextNode, _ = getPos(nextX, cur.y)
		}
	} else if cur.value == "L" {
		if cur.prev.y < cur.y { // if pipe came from the north.
			nextNode, _ = getPos(cur.x+1, cur.y)
		} else { // if pipe came from the east.
			nextNode, _ = getPos(cur.x, cur.y-1)
		}
	} else if cur.value == "J" {
		if cur.prev.y < cur.y { // if pipe came from the north.
			nextNode, _ = getPos(cur.x-1, cur.y)
		} else { // if pipe came from the east.
			nextNode, _ = getPos(cur.x, cur.y-1)
		}
	} else if cur.value == "7" {
		if cur.prev.x < cur.x { // if pipe came from the west.
			nextNode, _ = getPos(cur.x, cur.y+1)
		} else { // if pipe came from the south.
			nextNode, _ = getPos(cur.x+1, cur.y)
		}
	} else if cur.value == "F" {
		if cur.prev.y > cur.y { // if pipe came from the north.
			nextNode, _ = getPos(cur.x+1, cur.y)
		} else { // if pipe came from the east.
			nextNode, _ = getPos(cur.x, cur.y+1)
		}
	}

	// L, J, 7, F

	if nextNode != nil {
		cur.next = nextNode
		nextNode.prev = &cur
		return *nextNode, nil
	}

	return cur, errors.New("No position found")
}
