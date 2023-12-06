package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	//count := 0
	scanner := bufio.NewScanner(file)

	step := 0
	sets := make(map[int][][]int)

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
			sets[step] = append(sets[step], entries)
		}
	}

	sum := 0
	for seed := 0; seed < len(seeds); seed += 2 {
		for span := seeds[seed]; span < seeds[seed]+seeds[seed+1]; span++ {
			value := span
			for set := 0; set < len(sets); set++ {
				nodes := sets[set]
				for _, node := range nodes {
					if value >= node[1] && value < node[1]+node[2] {
						value = value - node[1] + node[0]
						break
					}
				}
			}

			if value < sum || sum == 0 {
				sum = value
			}
		}
	}

	fmt.Println("The answer is: `" + fmt.Sprint(sum) + "`")

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}
