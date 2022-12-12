package main

import (
	"bufio"
	"fmt"
	"os"
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
		line := scanner.Text()
		bag1 := line[0 : len(line)/2]
		bag2 := line[len(line)/2:]

	out:
		for i := 0; i < len(bag1); i++ {
			for j := 0; j < len(bag2); j++ {
				if bag1[i] == bag2[j] {
					if bag1[i] < 97 {
						sum += (int(bag1[i]) - 38)
					} else {
						sum += (int(bag1[i]) - 96)
					}
					break out
				}
			}
		}
	}

	fmt.Printf("%v", sum)
}
