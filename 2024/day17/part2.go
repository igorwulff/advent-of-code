package main

import (
	"fmt"
	"reflect"

	"github.com/igorwulff/advent-of-code/2024/day17/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	comp := shared.ParseInput(input)

	a := 1
	for {
		out := []int{}
		comp.Pointer = 0
		comp.Registers["A"] = a
		for comp.Pointer < len(comp.Program) {
			instruction := comp.Program[comp.Pointer]
			operand := comp.Program[comp.Pointer+1]
			out = append(out, comp.Instruction(instruction, operand)...)
		}

		match := true
		for i := 1; i <= len(out); i++ {
			if out[len(out)-i] != comp.Program[len(comp.Program)-i] {
				match = false
			}
		}

		if match {
			if reflect.DeepEqual(out, comp.Program) {
				break
			}

			a *= 8
		} else {
			a++
		}
	}

	return fmt.Sprint(a)
}
