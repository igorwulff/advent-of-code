package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	r, _ := regexp.Compile("([0-9A-Z-]+)")
	seq := 0
	for scanner.Scan() {
		raw := r.FindAllString(scanner.Text(), -1)
		data := make([]int, len(raw))

		for k, v := range raw {
			value, _ := strconv.Atoi(v)
			data[k] = value
		}

		h := history(data)
		seq += h
	}

	fmt.Println("")
	log.Printf("Execution time: %s", time.Since(start))
	fmt.Printf("Next sequence: %d", seq)
	fmt.Println("")
}

func history(data []int) int {
	diff := make([]int, len(data)-1)

	for i := 0; i < len(data)-1; i++ {
		diff[i] = calcDiff(data[i], data[i+1])
	}

	if len(diff) > 1 {
		return history(diff) + data[0]
	}

	return data[0]
}

func calcDiff(a, b int) int {
	return a - b
}
