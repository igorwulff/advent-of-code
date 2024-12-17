package main

import (
	"fmt"
	"strings"

	"github.com/igorwulff/advent-of-code/2024/day17/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	comp := shared.ParseInput(input)

	out := []int{}
	for comp.Pointer < len(comp.Program) {
		instruction := comp.Program[comp.Pointer]
		operand := comp.Program[comp.Pointer+1]

		out = append(out, comp.Instruction(instruction, operand)...)
	}

	return strings.Trim(
		strings.Join(
			strings.Fields(fmt.Sprint(out)), ",",
		), "[]",
	)
}
