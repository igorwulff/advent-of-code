package shared

import (
	"sort"
	"strconv"
	"strings"

	"github.com/igorwulff/advent-of-code/utils"
)

var ConnectCount = 1000

type Circuit struct {
	Junctions []*Junction
}

type Junction struct {
	X, Y, Z int
	Circuit *Circuit
}

type Distance struct {
	A    *Junction
	B    *Junction
	Dist int
}

func FindDistances(junctions []*Junction) []Distance {
	distances := make([]Distance, 0, len(junctions)*len(junctions))

	for i, a := range junctions {
		for j, b := range junctions {
			if a == b || j < i {
				continue
			}

			distances = append(distances, Distance{
				A:    a,
				B:    b,
				Dist: utils.PowInt(utils.AbsInt(a.X-b.X), 2) + utils.PowInt(utils.AbsInt(a.Y-b.Y), 2) + utils.PowInt(utils.AbsInt(a.Z-b.Z), 2),
			})
		}
	}

	// Sort by Dist value... so we can easily traverse it.
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Dist < distances[j].Dist
	})

	return distances
}

func ParseInput(input string) []Distance {
	junctions := make([]*Junction, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		j := Junction{
			X: x,
			Y: y,
			Z: z,
		}

		junctions = append(junctions, &j)
	}

	return FindDistances(junctions)
}
