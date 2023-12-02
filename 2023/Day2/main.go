package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		game, _ := strconv.Atoi(strings.Split(line[0], "Game ")[1])
		sets := strings.Split(line[1], ";")

		for _, set := range sets {
			parts := strings.Split(set, ",")

			blue := 0
			red := 0
			green := 0
			for _, part := range parts {
				pair := strings.Split(strings.TrimSpace(part), " ")

				if strings.Contains(pair[1], "blue") {
					color, _ := strconv.Atoi(pair[0])
					blue += color
				} else if strings.Contains(pair[1], "red") {
					color, _ := strconv.Atoi(pair[0])
					red += color
				} else if strings.Contains(pair[1], "green") {
					color, _ := strconv.Atoi(pair[0])
					green += color
				}
			}

			if blue > 14 || red > 12 || green > 13 {
				game = 0
			}
		}

		fmt.Println(game)
		sum += game
	}

	fmt.Println(sum)
}
