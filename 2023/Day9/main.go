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

type node struct {
	value      string
	left       *node
	leftValue  string
	right      *node
	rightValue string
}

func main() {
	start := time.Now()

	file, err := os.Open("./sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	r, _ := regexp.Compile("([0-9A-Z]+)")

	sum := 0
	scanner.Scan()
	raw := r.FindAllString(scanner.Text(), -1)
	data := make([]int, len(raw))
	for _, v := range raw {
		value, _ := strconv.Atoi(v)
		if value > 0 {
			data = append(data, value)
		}
	}
	fmt.Println(data)

	for i := 0; i < len(data)-1; i++ {
		diff := data[i+1] - data[i]
		fmt.Println(diff)
	}

	log.Printf("Execution time: %s", time.Since(start))
	fmt.Printf("Steps needed: %d", sum)
}
