package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer file.Close()

	scanner := bufio.NewScanner(file)

	solution(scanner)
}

func solution(scanner *bufio.Scanner) {
	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		m := strings.Split(line, " ")
		nums := stringToIntSlice(m)
		lastVals := [][]int{}
		fl := []int{nums[0], nums[len(nums)-1]}
		// Storing the first and last values for each generations
		lastVals = append(lastVals, fl)
		for {
			newNums := []int{}
			allZero := true
			for i := 0; i < len(nums)-1; i++ {
				a := nums[i]
				b := nums[i+1]
				diff := b - a
				if diff != 0 {
					allZero = false
				}
				newNums = append(newNums, diff)
			}
			nums = newNums
			fl := []int{nums[0], nums[len(nums)-1]}
			lastVals = append(lastVals, fl)
			if allZero {
				break
			}

		}

		firstVal := 0
		lastVal := 0
		for i := len(lastVals) - 1; i >= 0; i-- {
			previousVals := lastVals[i]
			// FOR PART 1
			previousFirst := previousVals[0]
			firstVal = previousFirst - firstVal
			// FOR PART 2
			previousLast := previousVals[1]
			lastVal = lastVal + previousLast
		}

		// FOR PART 1
		sum1 += lastVal

		// FOR PART 2
		sum2 += firstVal
	}

	fmt.Println("Part 1", sum1)
	fmt.Println("Part 1", sum2)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertToInt(s string) int {
	s2n := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
	if s2n[s] != 0 || s == "0" {
		return s2n[s]
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return num
}

func readFile(fname string) string {
	data, err := os.ReadFile(fname)
	check(err)
	return string(data)
}

func readLines(fname string) []string {
	data := readFile(fname)
	lines := strings.Split(data, "\n")
	// popping off the blank line
	return lines[:len(lines)-1]
}

func stringTwoSplit(s string, sep string) (string, string) {
	split := strings.Split(s, sep)
	if len(split) < 2 {
		return "", ""
	}
	return split[0], split[1]
}

func reMatch(pattern string, s string) ([]string, [][]int, [][]string) {
	r, _ := regexp.Compile(pattern)
	return r.FindAllString(s, -1), r.FindAllStringIndex(s, -1), r.FindAllStringSubmatch(s, -1)
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func getMapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func getMapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for k := range m {
		v := m[k]
		values = append(values, v)
	}
	return values
}

func stringToIntSlice(s []string) []int {
	var iSlice []int

	for _, v := range s {
		cleanVal := strings.TrimSpace(v)
		if cleanVal == "" {
			continue
		}
		iSlice = append(iSlice, convertToInt(cleanVal))
	}

	return iSlice
}
