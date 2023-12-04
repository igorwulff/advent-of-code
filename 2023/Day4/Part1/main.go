package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("sdf")
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		line = strings.Split(line[1], "|")

		winners := strings.Split(line[0], " ")
		losers := strings.Split(line[1], " ")

		count := 0
		for _, v1 := range winners {
			if v1 != "" {
				for _, v2 := range losers {
					if v2 != "" {
						if v1 == v2 {
							if count == 0 {
								count = 1
							} else {
								count *= 2
							}
						}
					}
				}
			}
		}

		sum += count

	}

	fmt.Println("The answer is: `" + fmt.Sprint(sum) + "`")
}
