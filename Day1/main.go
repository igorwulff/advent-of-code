package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := make([]int, 5)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			sum = append(sum, count)
			count = 0
		} else {
			intValue, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			count += intValue
		}
	}

	sort.Ints(sum)
	sumLength := len(sum)

	fmt.Printf("%v", sum[sumLength-1]+sum[sumLength-2]+sum[sumLength-3])
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
