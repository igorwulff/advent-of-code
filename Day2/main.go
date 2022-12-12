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

	for scanner.Scan() {
		result := strings.Split(scanner.Text(), " ")

		switch result[1] {
		case "X": // Rock X or A
			points += 1
		case "Y": // Paper Y or B
			points += 2
		case "Z": // Scizzors Z or C
			points += 3
		}

		if (result[0] == "A" && result[1] == "X") || (result[0] == "B" && result[1] == "Y") || (result[0] == "C" && result[1] == "Z") {
			points += 3
		} else if (result[0] == "A" && result[1] == "Y") || (result[0] == "B" && result[1] == "Z") || (result[0] == "C" && result[1] == "X") {
			points += 6
		}
	}

	fmt.Printf("%v", points)
}
