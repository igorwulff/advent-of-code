package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Exported function to be called by the main application
func Part1(input string) string {
	values := strings.Split(input, " ")

	sum := 0
	for _, value := range values {
		sum += Blink(25, value)
	}

	return fmt.Sprint(sum)
}

func Blink(i int, value string) int {
	if i == 0 {
		return 1
	}

	i--

	if value == "0" {
		return Blink(i, "1")
	}

	if len(value)%2 == 0 {
		l, _ := strconv.Atoi(strings.TrimLeft(value[:len(value)/2], "0"))
		r, _ := strconv.Atoi(strings.TrimLeft(value[len(value)/2:], "0"))

		return Blink(i, strconv.Itoa(l)) + Blink(i, strconv.Itoa(r))
	}

	number, _ := strconv.Atoi(value)
	return Blink(i, strconv.Itoa(number*2024))
}
