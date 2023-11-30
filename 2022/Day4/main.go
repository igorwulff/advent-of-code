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

	out:
		for i := result1[0]; i <= result1[1]; i++ {
			for j := result2[0]; j <= result2[1]; j++ {
				if i == j {
					sum += 1
					break out
				}
			}
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
