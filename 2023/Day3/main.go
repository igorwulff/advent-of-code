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

// 65684750
// NOT CORRECT: 65299218
// 68864416
// 74528807
func main() {
	lines, _ := readLines("input.txt")

	prevLine := ""
	curLine := lines[0]
	nextLine := lines[1]

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
			if string(letter) == "*" {
				list := make([]string, 0, 5)
				start := index - 1
				if start < 0 {
					start = 0
				}

				list = lookup(prevLine, list, start)
				list = lookup(curLine, list, start)
				list = lookup(nextLine, list, start)

				if len(list) == 2 {
					left, _ := strconv.Atoi(list[0])
					right, _ := strconv.Atoi(list[1])
					sum += (left * right)
				}
			}
		}
	}

	fmt.Println("The answer is: `" + fmt.Sprint(sum) + "`")
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

func lookup(line string, list []string, start int) []string {
	if line != "" {
		value := ""
		for key, sign := range line[start:] {
			if unicode.IsDigit(sign) {
				if value == "" && key == 0 {
					lookup := start - 1
					for lookup >= 0 {
						_, error := strconv.Atoi(string(line[lookup]))
						if error == nil {
							value = string(line[lookup]) + value
							lookup--
						} else {
							break
						}
					}
				}

				value += string(sign)
				if start+key+1 == len(line) {
					list = append(list, value)
				}
			} else if value != "" {
				list = append(list, value)
				value = ""
				if string(sign) != "*" && key > 1 {
					break
				}
			} else if value == "" && key > 1 {
				break
			}
		}
	}

	return list
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
