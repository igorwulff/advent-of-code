package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	points := 0
	shape := ""

	for scanner.Scan() {
		result := strings.Split(scanner.Text(), " ")

		// X = lose, Y = draw, Z = win
		if result[1] == "Y" {
			shape = result[0]
			points += 3
		} else if result[1] == "X" {
			if result[0] == "A" {
				shape = "C"
			} else if result[0] == "B" {
				shape = "A"
			} else {
				shape = "B"
			}
		} else if result[1] == "Z" {
			if result[0] == "A" {
				shape = "B"
			} else if result[0] == "B" {
				shape = "C"
			} else {
				shape = "A"
			}
			points += 6
		}

		switch shape {
		case "A": // Rock X or A
			points += 1
		case "B": // Paper Y or B
			points += 2
		case "C": // Scizzors Z or C
			points += 3
		}
	}

	fmt.Printf("%v", points)
}
