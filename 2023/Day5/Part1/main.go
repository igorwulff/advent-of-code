package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//count := 0
	scanner := bufio.NewScanner(file)

	step := 0
	//set := make([][]int, 10)
	set := make(map[int][][]int)

	var data []string
	if scanner.Scan() {
		// Map string of seeds to seeds slice
		data = strings.Split(scanner.Text(), ": ")
		data = strings.Split(data[1], " ")

		// Skip next line.
		scanner.Scan()
		scanner.Scan()
	}

	seeds := make([]int, len(data))
	for k, i := range data {
		seed, _ := strconv.Atoi(i)
		seeds[k] = seed
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			step++
			// Skip next line.
			scanner.Scan()

		} else {
			var entries []int
			data := strings.Split(scanner.Text(), " ")
			for _, i := range data {
				entry, _ := strconv.Atoi(i)
				entries = append(entries, entry)
			}
			set[step] = append(set[step], entries)
		}
	}
	fmt.Println("==========")
	sum := 0
	for _, seed := range seeds {
		var value = seed
		fmt.Print("Seed:")
		fmt.Println(seed)
		for i := 0; i < len(set); i++ {
			nodes := set[i]
			fmt.Println(nodes)
			for _, node := range nodes {
				if value >= node[1] && value < node[1]+node[2] {
					fmt.Println(value - node[1] + node[0])
					value = value - node[1] + node[0]
					break
				}
			}
			fmt.Println(value)
		}

		fmt.Print("Location: ")
		fmt.Println(value)

		if value < sum || sum == 0 {
			sum = value
		}
	}

	fmt.Println(sum)

}
