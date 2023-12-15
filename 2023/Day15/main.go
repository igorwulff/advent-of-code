package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	scanner := bufio.NewScanner(file)
	sum := 0

	list := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		list = append(list, strings.Split(text, ",")...)
	}

	for _, sequence := range list {
		value := 0
		for _, letter := range sequence {
			value += int(letter)
			value *= 17
			value %= 256
		}
		sum += value
	}

	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}
