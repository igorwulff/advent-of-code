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

	r, _ := regexp.Compile("([1-9A-Z]+)")
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

	startNodes := make([]*node, 0)
	for key, node := range nodes {
		if string(key[2]) != "A" {
			continue
		}
		startNodes = append(startNodes, node)
	}

	steps := 0
out:
	for {
		for _, direction := range coord {
			steps++
			newNodes := make([]*node, 0)
			for _, node := range startNodes {
				if string(direction) == "L" {
					newNodes = append(newNodes, node.left)
				} else {
					newNodes = append(newNodes, node.right)
				}
			}
			startNodes = newNodes

			matches := 0

			//fmt.Printf("Run %d", steps)
			for _, node := range startNodes {
				if string(node.value[2]) == "Z" {
					matches++
				}
			}

			if matches == len(startNodes) {
				break out
			}
		}
	}

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
	fmt.Printf("Steps needed: %d", steps)

	/// 21883 to low answer.
}
