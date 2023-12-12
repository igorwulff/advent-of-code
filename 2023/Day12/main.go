package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	r, _ := regexp.Compile("(#+)")

	for scanner.Scan() {
		row := scanner.Text()
		record := strings.Split(row, " ")
		conditions := record[0]
		groups := strings.Split(record[1], ",")

		matches := providePossibleMatches(conditions, 0, "")
		combinations := 0
		for _, match := range matches {
			text := r.FindAllString(match, -1)

			if len(text) == len(groups) {
				isMatch := true
				for i := 0; i < len(text); i++ {
					groupInt, _ := strconv.Atoi(groups[i])
					if len(text[i]) != groupInt {
						isMatch = false
					}
				}

				if isMatch {
					combinations++
				}
			}
		}
		sum += combinations
	}

	fmt.Printf("Combinations Distances: %d", sum)
	fmt.Println()
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func providePossibleMatches(input string, position int, buildup string) []string {
	matches := make([]string, 0)

	if len(input) == position {
		return append(matches, buildup)
	}

	if string(input[position]) == "?" {
		matches = append(matches, providePossibleMatches(input, position+1, buildup+".")...)
		matches = append(matches, providePossibleMatches(input, position+1, buildup+"#")...)
	} else {
		matches = append(matches, providePossibleMatches(input, position+1, buildup+string(input[position]))...)
	}
	return matches
}
