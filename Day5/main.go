package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stackList := make([]Stack, 9)
	stackList[0] = append(stackList[0], "D", "T", "R", "B", "J", "L", "W", "G")
	stackList[1] = append(stackList[1], "S", "W", "C")
	stackList[2] = append(stackList[2], "R", "Z", "T", "M")
	stackList[3] = append(stackList[3], "D", "T", "C", "H", "S", "P", "V")
	stackList[4] = append(stackList[4], "G", "P", "T", "L", "D", "Z")
	stackList[5] = append(stackList[5], "F", "B", "R", "Z", "J", "Q", "C", "D")
	stackList[6] = append(stackList[6], "S", "B", "D", "J", "M", "F", "T", "R")
	stackList[7] = append(stackList[7], "L", "H", "R", "B", "T", "V", "M")
	stackList[8] = append(stackList[8], "Q", "P", "D", "S", "V")

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		count, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		indexFrom, err := strconv.Atoi(line[3])
		if err != nil {
			panic(err)
		}
		indexFrom -= 1

		indexTo, err := strconv.Atoi(line[5])
		if err != nil {
			panic(err)
		}
		indexTo -= 1

		for i := 0; i < count; i++ {
			stackItem := stackList[indexFrom].Pop()
			stackList[indexTo].Push(stackItem)
		}
	}

	for _, letter := range stackList {
		fmt.Print(letter.Pop())
		//fmt.Printf("%s", letter.Pop())
	}
}

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element
	}
}
