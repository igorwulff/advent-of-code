package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		first := ""
		last := ""
		word := ""
		for _, v := range line {
			if unicode.IsDigit(v) {
				if first == "" {
					first = string(v)
					last = string(v)
				} else {
					last = string(v)
				}
			} else {
				word += string(v)
				for digit, text := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
					if strings.Contains(word, text) {
						if first == "" {
							first = fmt.Sprint(digit + 1)
							last = fmt.Sprint(digit + 1)
						} else {
							last = fmt.Sprint(digit + 1)
						}
						word = "" + word[len(word)-1:]
					}
				}
			}
		}

		count, _ := strconv.Atoi(first + last)
		sum += count
	}

	fmt.Println("The answer is: `" + fmt.Sprint(sum) + "`")
}
