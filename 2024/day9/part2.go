package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day9/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	disk := shared.ParseInput(input)

	fmt.Println("LENGHT:", len(disk))

	curDir := disk[len(disk)-1].Id
	for dir := curDir; dir >= 0; dir-- {
		for i := 0; i < len(disk); i++ {
			if disk[i] == nil {
				continue
			}

			if disk[i].Id != dir {
				continue
			}

			idx := 0
			size := 0
			for j := 0; j <= i; j++ {
				if disk[j] != nil {
					idx = 0
					size = 0
					continue
				} else if idx == 0 {
					idx = j
				}

				size++
				if size == disk[i].Size {

					for k := 0; k < size; k++ {
						disk[idx+k] = disk[i+k]
						disk[i+k] = nil
					}
					break
				}
			}
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
