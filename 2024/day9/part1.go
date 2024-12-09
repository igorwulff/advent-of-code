package main

import (
	"fmt"
	"strconv"
	"strings"
)

type File struct {
	Id int
}

// Exported function to be called by the main application
func Part1(input string) string {
	disk := make([]*File, 0)

	lines := strings.Split(input, "\n")
	id := 0
	for _, line := range lines {
		for k, c := range line {
			l, _ := strconv.Atoi(string(c))
			if k%2 == 0 {
				file := &File{
					Id: id,
				}
				id++

				for i := 0; i < l; i++ {
					disk = append(disk, file)
				}
			} else {
				for i := 0; i < l; i++ {
					disk = append(disk, nil)
				}
			}
		}
	}

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
