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
	coordLength := len(coord);
out:
	for {
		for i:=0; i<coordLength; i++ {
			steps++
			if coord[i] == 76 { // L
				startNodes[0] = startNodes[0].left
				startNodes[1] = startNodes[1].left
				startNodes[2] = startNodes[2].left
				startNodes[3] = startNodes[3].left
				startNodes[4] = startNodes[4].left
				startNodes[5] = startNodes[5].left
			} else {
				startNodes[0] = startNodes[0].right
				startNodes[1] = startNodes[1].right
				startNodes[2] = startNodes[2].right
				startNodes[3] = startNodes[3].right
				startNodes[4] = startNodes[4].right
				startNodes[5] = startNodes[5].right
			}

			if (startNodes[0].value[2] == 90 && startNodes[1].value[2] == 90 && startNodes[2].value[2] == 90 && startNodes[3].value[2] == 90 && startNodes[4].value[2] == 90 && startNodes[5].value[2] == 90) {
				break out
			}
		}

		fmt.Println(steps);
	}


	log.Printf("Execution time: %s", time.Since(start))
	fmt.Printf("Steps needed: %d", steps)

	/// 21883 to low answer.
}
