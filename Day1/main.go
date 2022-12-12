package main

import (
	"bufio"
	"fmt"
	"os"
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

	highestCalories := 0
	for _, v := range sum {
		highestCalories = max(highestCalories, v)
	}

	fmt.Printf("%v", highestCalories)
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
