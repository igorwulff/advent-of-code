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

	time := make([]int, 4)
	distance := make([]int, 4)

	scanner.Scan()
	data := strings.Split(scanner.Text(), ":")
	data = strings.Split(strings.Trim(data[1], " "), " ")
	i := 0
	for _, v := range data {
		v = strings.Trim(v, " ")
		value, _ := strconv.Atoi(v)
		if value > 0 {
			time[i] = value
			i++
		}
	}

	scanner.Scan()
	data = strings.Split(scanner.Text(), ":")
	data = strings.Split(data[1], " ")
	i = 0
	for _, v := range data {
		v = strings.Trim(v, " ")
		value, _ := strconv.Atoi(v)
		if value > 0 {
			distance[i] = value
			i++
		}
	}

	totalWins := 1
	for k, t := range time {
		d := distance[k]
		wins := 0
		for i := 1; i < t; i++ {
			if i*(time[k]-i) > d {
				wins++
			}
		}

		totalWins *= wins
	}

	sum := totalWins

	fmt.Println("The answer is: `" + fmt.Sprint(sum) + "`")
}
