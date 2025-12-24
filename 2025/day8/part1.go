package main

import (
	"fmt"
	"sort"

	"github.com/igorwulff/advent-of-code/2025/day8/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	distances := shared.ParseInput(input)

	circuits := make([]*shared.Circuit, 0)

	for i := 0; i < shared.ConnectCount; i++ {
		dist := distances[i]

		a := dist.A
		b := dist.B

		if a.Circuit == nil && b.Circuit == nil {
			c := &shared.Circuit{
				Junctions: make([]*shared.Junction, 0),
			}
			c.Junctions = append(c.Junctions, a, b)
			a.Circuit = c
			b.Circuit = c
			circuits = append(circuits, c)
		} else if a.Circuit != nil && b.Circuit == nil {
			a.Circuit.Junctions = append(a.Circuit.Junctions, b)
			b.Circuit = a.Circuit
		} else if a.Circuit == nil && b.Circuit != nil {
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

	sum := 1
	for i := range 3 {
		sum *= len(circuits[i].Junctions)
	}

	return fmt.Sprint(sum)
}

// 2665464 to high
