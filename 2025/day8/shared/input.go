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

func ConnectJunctions(distances []Distance, maxIterations int) ([]*Circuit, Distance) {
	circuits := make([]*Circuit, 0)
	var lastDistance Distance

	for i := 0; i < maxIterations; i++ {
		dist := distances[i]

		a := dist.A
		b := dist.B

		if a.Circuit == nil && b.Circuit == nil {
			lastDistance = dist
			c := &Circuit{
				Junctions: make([]*Junction, 0),
			}
			c.Junctions = append(c.Junctions, a, b)
			a.Circuit = c
			b.Circuit = c
			circuits = append(circuits, c)
		} else if a.Circuit != nil && b.Circuit == nil {
			lastDistance = dist
			a.Circuit.Junctions = append(a.Circuit.Junctions, b)
			b.Circuit = a.Circuit
		} else if a.Circuit == nil && b.Circuit != nil {
			lastDistance = dist
			b.Circuit.Junctions = append(b.Circuit.Junctions, a)
			a.Circuit = b.Circuit
		} else if a.Circuit != b.Circuit {
			ca, cb := a.Circuit, b.Circuit

			// Ensure ca is the larger circuit
			if len(ca.Junctions) < len(cb.Junctions) {
				ca, cb = cb, ca
			}

			// Merge cb into ca
			ca.Junctions = append(ca.Junctions, cb.Junctions...)
			for _, jn := range cb.Junctions {
				jn.Circuit = ca
			}

			// Remove cb from circuits slice
			for idx, c := range circuits {
				if c == cb {
					circuits = append(circuits[:idx], circuits[idx+1:]...)
					break
				}
			}
		}
	}

	// Sort by Dist value... so we can easily traverse it.
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i].Junctions) > len(circuits[j].Junctions)
	})

	return circuits, lastDistance
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
