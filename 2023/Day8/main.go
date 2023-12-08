package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

type node struct {
	value      string
	left       *node
	leftValue  string
	right      *node
	rightValue string
}

func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	r, _ := regexp.Compile("([A-Z]+)")
	nodes := make(map[string]*node)
	scanner.Scan()
	coord := scanner.Text()
	scanner.Scan()

	// AAA = (BBB, BBB)
	for scanner.Scan() {
		raw := r.FindAllString(scanner.Text(), 3)
		nodes[raw[0]] = &node{value: raw[0], leftValue: raw[1], rightValue: raw[2]}
	}

	for _, nodeI := range nodes {
		nodeI.left = nodes[nodeI.leftValue]
		nodeI.right = nodes[nodeI.rightValue]
	}

	curNode := nodes["AAA"]
	steps := 0
	for {
		for _, direction := range coord {
			steps++
			if string(direction) == "L" {
				curNode = curNode.left
			} else {
				curNode = curNode.right
			}

			if string(direction) == "ZZZ" {
				break
			}
		}
		if curNode.value == "ZZZ" {
			break
		}
	}

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
	fmt.Printf("Steps needed: %d", steps)
}
