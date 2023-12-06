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

	times := make([]int, 1)
	distance := make([]int, 1)

	scanner.Scan()
	data := strings.Split(scanner.Text(), ":")
	data = strings.Split(strings.Trim(data[1], " "), " ")
	value := ""
	for _, v := range data {
		value += strings.Trim(v, " ")
	}
	times[0], _ = strconv.Atoi(value)

	scanner.Scan()
	data = strings.Split(scanner.Text(), ":")
	data = strings.Split(strings.Trim(data[1], " "), " ")
	value = ""
	for _, v := range data {
		value += strings.Trim(v, " ")
	}
	distance[0], _ = strconv.Atoi(value)

	totalWins := 1
	for k, t := range times {
		d := distance[k]
		wins := 0
		for i := 1; i < t; i++ {
			if i*(times[k]-i) > d {
				wins++
			} else if wins > 0 && d > times[k]+i {
				break
			}
		}

		totalWins *= wins
	}

	sum := totalWins

	fmt.Println("The answer is: `" + fmt.Sprint(sum) + "`")

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}
