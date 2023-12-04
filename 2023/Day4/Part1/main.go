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

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
	}

	cards := make([]int, count)
	for k, _ := range cards {
		cards[k] = 1
	}

	sum := 0
	card := 1

	file, err = os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		line = strings.Split(line[1], "|")

		winners := strings.Split(line[0], " ")
		losers := strings.Split(line[1], " ")

		matches := 0
		for _, v1 := range winners {
			if v1 != "" {
				for _, v2 := range losers {
					if v2 != "" {
						if v1 == v2 {
							matches++
						}
					}
				}
			}
		}

		multiplier := cards[card-1]
		for i := card; i < matches+card && i < count; i++ {
			cards[i] += multiplier
		}
		card++
	}

	for _, v := range cards {
		sum += v
	}

	fmt.Println("The answer is: `" + fmt.Sprint(sum) + "`")
}
