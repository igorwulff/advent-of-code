package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stack := make([]Stack, 9)
	stack[0] = append(stack[0], "D", "T", "R", "B", "J", "L", "W", "G")
	stack[1] = append(stack[1], "S", "W", "C")
	stack[2] = append(stack[2], "R", "Z", "T", "M")
	stack[3] = append(stack[3], "D", "T", "C", "H", "S", "P", "V")
	stack[4] = append(stack[4], "G", "P", "T", "L", "D", "Z")
	stack[5] = append(stack[5], "F", "B", "R", "Z", "J", "Q", "C", "D")
	stack[6] = append(stack[6], "S", "B", "D", "J", "M", "F", "T", "R")
	stack[7] = append(stack[7], "L", "H", "R", "B", "T", "V", "M")
	stack[8] = append(stack[8], "Q", "P", "D", "S", "V")

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

		from, err := strconv.Atoi(line[3])
		if err != nil {
			panic(err)
		}
		from -= 1

		to, err := strconv.Atoi(line[5])
		if err != nil {
			panic(err)
		}
		to -= 1

		mover := make(Stack, count)
		copy(mover, stack[from][len(stack[from])-count:])
		stack[from] = stack[from][:len(stack[from])-count]
		stack[to] = append(stack[to], mover...)
	}

	for _, letter := range stack {
		fmt.Print(letter[len(letter)-1])
	}
}

type Stack []string

func (s *Stack) IsNotEmpty() bool {
	return len(*s) != 0
}

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

func (s *Stack) Unshift() string {
	if s.IsEmpty() {
		return ""
	} else {
		element := (*s)[0] // The first element is the one to be dequeued.
		*s = (*s)[1:]      // Slice off the element once it is dequeued.
		return element
	}
}
