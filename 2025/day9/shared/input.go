package shared

import (
	"sort"
	"strconv"
	"strings"

	"github.com/igorwulff/advent-of-code/utils"
)

type Corner struct {
	X, Y int
}

type Distance struct {
	A    *Corner
	B    *Corner
	Size int
}

func FindDistances(corners []*Corner) []Distance {
	distances := make([]Distance, 0, len(corners)*len(corners))

	for i, a := range corners {
		for j, b := range corners {
			if a == b || j < i {
				continue
			}

			diffX := utils.AbsInt(a.X-b.X) + 1
			diffY := utils.AbsInt(a.Y-b.Y) + 1

			distances = append(distances, Distance{
				A:    a,
				B:    b,
				Size: diffX * diffY,
			})
		}
	}

	// Sort by Dist value... so we can easily traverse it.
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Size > distances[j].Size
	})

	return distances
}

func ParseInput(input string) []*Corner {
	corners := make([]*Corner, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		j := Corner{
			X: x,
			Y: y,
		}

		corners = append(corners, &j)
	}

	return corners
}
