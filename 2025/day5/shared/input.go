package shared

import (
	"sort"
	"strconv"
	"strings"
)

type Ingredients struct {
	Avail []int64
	Fresh []string
}

func ParseInput(input string) *Ingredients {
	ingr := Ingredients{
		Avail: make([]int64, 0, 500),
		Fresh: make([]string, 0, 500),
	}

	lines := strings.Split(input, "\n")
	step1 := true
	for _, line := range lines {
		if line == "" {
			step1 = false
			continue
		}

		if step1 {
			ingr.Fresh = append(ingr.Fresh, line)
		} else {
			id, _ := strconv.ParseInt(line, 10, 64)
			ingr.Avail = append(ingr.Avail, id)
		}
	}

	return &ingr
}

func (ingr Ingredients) CheckIfExists(id int64) bool {
	for _, v := range ingr.Fresh {
		r := ingr.GetRange(v)

		if id >= r.Start && id <= r.End {
			return true
		}
	}

	return false
}

func (ingr Ingredients) GetRange(id string) IdRange {
	r := strings.Split(id, "-")

	start, _ := strconv.ParseInt(r[0], 10, 64)
	end, _ := strconv.ParseInt(r[1], 10, 64)

	return IdRange{
		Start: start,
		End:   end,
	}
}

type IdRange struct {
	Start int64
	End   int64
}

func (ingr Ingredients) CountRanges(fresh []string) int64 {
	ranges := make([]IdRange, 0, 500)
	for _, r := range fresh {
		ranges = append(ranges, ingr.GetRange(r))
	}

	// Sort by Start value... so we can easily traverse it.
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	var idx int64 = 0
	var count int64 = 0
	// Iterate over each sorted range and identify overlaps and handle them for each case.
	for _, r := range ranges {
		if idx == 0 {
			idx = r.End
			count += r.End - r.Start + 1
			continue
		}

		if idx >= r.End {
			continue
		}

		if idx >= r.Start {
			r.Start = idx + 1
		}

		idx = r.End
		count += r.End - r.Start + 1
	}

	return count
}
