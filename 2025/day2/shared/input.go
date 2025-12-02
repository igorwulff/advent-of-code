package shared

import (
	"strconv"
	"strings"
)

type Ids struct {
	Start int
	End   int
}

func ParseInput(input string) []Ids {
	data := strings.Split(input, ",")
	output := make([]Ids, 0, len(data))

	for _, ids := range data {
		split := strings.Split(ids, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		output = append(output, Ids{
			Start: start,
			End:   end,
		})
	}

	return output
}
