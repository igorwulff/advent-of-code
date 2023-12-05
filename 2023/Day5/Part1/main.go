package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//count := 0
	scanner := bufio.NewScanner(file)
	
	step := 0
	map := make([][]string, 10)
	for i := range map {
		map[i] = make([]string, 3)
	}

	if scanner.Scan() {
		// Map string of seeds to seeds slice
		data := strings.Split(scanner.Text(), ":")
		data = strings.Split(data[1], " ")
		seeds := make([]string, len(data))
		seeds = append(seeds, data...)

		// Skip next line.
		scanner.Scan()
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			step++
			// Skip next line.
			scanner.Scan()
		} else {
			data := strings.Split(, scanner.Text())
			
		}
	}

	sum := 0

	fmt.Println("The answer is: `" + fmt.Sprint(sum) + "`")
}
