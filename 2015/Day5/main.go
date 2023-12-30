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

	var sum1, sum2 int

	vowels := map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true}
	forbidden := map[string]bool{"ab": true, "cd": true, "pq": true, "xy": true}

	for scanner.Scan() {
		text := scanner.Text()

		pairs := make([]string, 0)
		for i := 0; i < len(text)-1; i++ {
			pairs = append(pairs, string(text[i])+string(text[i+1]))
		}

		if isValidPart1(text, vowels, forbidden) {
			sum1++
		}

		if isValidPart2(text, pairs) {
			sum2++
		}
	}

	fmt.Println("Part1: " + fmt.Sprint(sum1))
	fmt.Println("Part2: " + fmt.Sprint(sum2))
	log.Printf("Execution time: %s", time.Since(start))
}

func isValidPart2(text string, pairs []string) bool {
	valid := false

out:
	for k, v := range pairs {
		for k2, v2 := range pairs {
			if k >= k2-1 && k <= k2+1 {
				continue
			}

			if v == v2 {
				valid = true
				continue out
			}
		}
	}

	if valid {
		for i := 0; i < len(text)-2; i++ {
			if string(text[i]) == string(text[i+2]) {
				return true
			}
		}
	}

	return false
}

func isValidPart1(text string, vowels map[string]bool, forbidden map[string]bool) bool {
	valid := true
	vowelCount := 0
	for f := range forbidden {
		if strings.Contains(text, f) {
			valid = false
			break
		}
	}

	if !valid {
		return false
	}

	valid = false
	for v, k := range text {
		for vowel := range vowels {
			if vowel == string(k) {
				vowelCount++
			}
		}

		if v+1 != len(text) && text[v+1] == text[v] {
			valid = true
		}
	}

	if !valid || vowelCount < 3 {
		return false
	}

	return true
}
