package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		result := strings.Split(scanner.Text(), ",")
		result1 := StringArrayToInt(strings.Split(result[0], "-"))
		result2 := StringArrayToInt(strings.Split(result[1], "-"))

		if result1[0] <= result2[0] && result1[1] >= result2[1] {
			sum += 1
		} else if result2[0] <= result1[0] && result2[1] >= result1[1] {
			sum += 1
		}
	}

	fmt.Print(sum)
}

func StringArrayToInt(stringArray []string) []int {
	intArray := make([]int, 0, len(stringArray))

	for _, value := range stringArray {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		intArray = append(intArray, intValue)
	}

	return intArray
}
