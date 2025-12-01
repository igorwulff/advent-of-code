package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2025/day1/shared"
)

func Part2(input string) string {
	dirs, steps := shared.ParseInput(input)
	counter := 0
	pos := 50

	for i, dir := range dirs {
		step := steps[i]

		switch dir {

		case "L":
			counter += step / 100
			step %= 100

			/*pos -= step
			if pos == 0 {
				counter++
			}
			if pos < 0 {
				counter++
				pos += 99
			}*/

			for range step {
				pos--
				if pos == 0 {
					counter++
				}

				if pos < 0 {
					pos = 99
				}
			}
		case "R":
			pos += step
			counter += pos / 100
			pos %= 100
		}

		fmt.Printf("Dial at %d = %s%d: %d\n", pos, dir, step, counter)
	}

	return fmt.Sprint(counter)
}

// 2288 is to low.
// 5984 is to high.
// 5978 is correct.
