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

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	text := ""
	for scanner.Scan() {
		text += scanner.Text()
	}
	answer := strings.Count(text, "(") - strings.Count(text, ")")
	fmt.Println("Part1: " + fmt.Sprint(answer))

	pos := 1
	for i := 0; i < len(text); i++ {
		if text[i] == '(' {
			pos++
		} else if text[i] == ')' {
			pos--
		}

		if pos == -1 {
			fmt.Println("Part2: " + fmt.Sprint(i))
			break
		}
	}

	log.Printf("Execution time: %s", time.Since(start))
}
