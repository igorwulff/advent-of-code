package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Cell struct {
	value string
}

func main() {
	lines, _ := readLines("input.txt")

	prevLine := ""
	curLine := lines[0]
	nextLine := lines[1]

	value := ""
	sum := 0
	for row, line := range lines {
		if row > 0 {
			prevLine = curLine
		}
		curLine = line
		if len(lines) > row+1 {
			nextLine = lines[row+1]
		} else {
			nextLine = ""
		}

		for index, letter := range line {
			if unicode.IsDigit(letter) {
				value += string(letter)
			}

			if (!unicode.IsDigit(letter) && value != "") || len(curLine)-1 == index {

				start := index - len(value) - 1
				end := index + 1

				if len(curLine) < end {
					end = len(curLine)
				}

				if start < 0 {
					start = 0
				}

				if prevLine != "" {
					for _, sign := range prevLine[start:end] {
						if string(sign) != "." && !unicode.IsDigit(sign) {
							digit, _ := strconv.Atoi(value)
							sum += digit
							fmt.Println("========================")
							fmt.Println(value)
							fmt.Println(string(sign))
							fmt.Println("========================")
							value = ""
							continue

						}
					}
				}

				for _, sign := range curLine[start:end] {
					if string(sign) != "." && !unicode.IsDigit(sign) {
						digit, _ := strconv.Atoi(value)
						sum += digit
						fmt.Println("========================")
						fmt.Println(value)
						fmt.Println(string(sign))
						fmt.Println("========================")
						value = ""
						continue

					}
				}

				if nextLine != "" {
					for _, sign := range nextLine[start:end] {
						if string(sign) != "." && !unicode.IsDigit(sign) {
							digit, _ := strconv.Atoi(value)
							sum += digit
							fmt.Println("========================")
							fmt.Println(value)
							fmt.Println(string(sign))
							fmt.Println("========================")
							value = ""
							continue

						}
					}
				}

				value = ""
			}
		}
	}

	fmt.Println("ANSWER:")
	fmt.Println(sum)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// NOTE: this isn't multi-Unicode-codepoint aware, like specifying skintone or
//
//	gender of an emoji: https://unicode.org/emoji/charts/full-emoji-modifiers.html
func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
