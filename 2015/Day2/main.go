package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	r, _ := regexp.Compile("([0-9]+)")

	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		part := r.FindAllString(scanner.Text(), -1)

		l, _ := strconv.Atoi(part[0])
		w, _ := strconv.Atoi(part[1])
		h, _ := strconv.Atoi(part[2])

		sum1 += 2 * l * w
		sum1 += 2 * w * h
		sum1 += 2 * h * l
		sum1 += min(l*w, w*h, h*l)

		sum2 += min(l+l+w+w, l+l+h+h, w+w+h+h) + (l * w * h)
	}

	fmt.Println("Part1: " + fmt.Sprint(sum1))
	fmt.Println("Part2: " + fmt.Sprint(sum2))
	log.Printf("Execution time: %s", time.Since(start))
}
