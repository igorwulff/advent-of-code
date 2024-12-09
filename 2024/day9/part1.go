package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day9/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	disk := shared.ParseInput(input)

	j := len(disk)
	for i := 0; i < j; i++ {
		if disk[i] != nil {
			continue
		}

		for j := len(disk) - 1; j > i; j-- {
			if disk[j] == nil {
				continue
			}

			disk[i] = disk[j]
			disk[j] = nil
			break
		}
	}

	sum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == nil {
			continue
		}

		sum += disk[i].Id * i
	}

	return fmt.Sprint(sum)
}
