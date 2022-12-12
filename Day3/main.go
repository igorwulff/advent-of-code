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
		bag1 := scanner.Text()
		scanner.Scan()
		bag2 := scanner.Text()
		scanner.Scan()
		bag3 := scanner.Text()

	out:
		for x := 0; x < len(bag1); x++ {
			for y := 0; y < len(bag2); y++ {
				for z := 0; z < len(bag3); z++ {
					if bag1[x] == bag2[y] && bag1[x] == bag3[z] {
						if bag1[x] < 97 {
							sum += (int(bag1[x]) - 38)
						} else {
							sum += (int(bag1[x]) - 96)
						}

						break out
					}
				}
			}
		}
	}

	fmt.Printf("%v", sum)
}
