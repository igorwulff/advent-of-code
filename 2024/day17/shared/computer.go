package shared

import (
	"github.com/igorwulff/advent-of-code/utils"
)

type Computer struct {
	Registers map[string]int // A, B, C
	Pointer   int            // Instruction Pointer
	Program   []int
}

func (c *Computer) ComboOperand(value int) int {
	switch value {
	case 0, 1, 2, 3:
		return value
	case 4:
		return c.Registers["A"]
	case 5:
		return c.Registers["B"]
	case 6:
		return c.Registers["C"]
	}

	return 0
}

func (c *Computer) Instruction(opcode int, operand int) []int {
	out := make([]int, 0)

	switch opcode {
	case 0: // adv
		c.Registers["A"] = c.Registers["A"] / utils.PowInt(2, c.ComboOperand(operand))

	case 1: // bxl
		c.Registers["B"] = c.Registers["B"] ^ operand

	case 2: // bst
		c.Registers["B"] = c.ComboOperand(operand) % 8

	case 3: // jnz
		if c.Registers["A"] != 0 {
			if c.Pointer != operand {
				c.Pointer = operand
				return out // Do not increment pointer
			}
		}

	case 4: // bxc
		c.Registers["B"] = c.Registers["B"] ^ c.Registers["C"]

	case 5: // out
		out = append(out, c.ComboOperand(operand)%8)

	case 6: // bdv
		c.Registers["B"] = c.Registers["A"] / utils.PowInt(2, c.ComboOperand(operand))

	case 7: // cdv
		c.Registers["C"] = c.Registers["A"] / utils.PowInt(2, c.ComboOperand(operand))
	}

	c.Pointer += 2

	return out
}
