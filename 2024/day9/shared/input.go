package shared

import (
	"strconv"
	"strings"
)

func ParseInput(input string) []*File {
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

	return disk
}
