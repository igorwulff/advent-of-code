package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type grid struct {
	cells map[string]int
}

func (g grid) addPos(x, y int) {
	g.cells[fmt.Sprint(x)+":"+fmt.Sprint(y)]++
}

func main() {
	start := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum1, sum2 int
	g := grid{cells: make(map[string]int)}
	g2 := grid{cells: make(map[string]int)}
	var x, y int
	var xr, yr, xs, ys int

	for scanner.Scan() {
		g.addPos(x, y)
		g2.addPos(x, y)

		text := scanner.Text()
		for i := 0; i < len(text); i++ {
			switch text[i] {
			case '>':
				x++
				if i%2 == 0 {
					xs++
				} else {
					xr++
				}
			case '<':
				x--
				if i%2 == 0 {
					xs--
				} else {
					xr--
				}
			case '^':
				y--
				if i%2 == 0 {
					ys++
				} else {
					yr++
				}
			case 'v':
				y++
				if i%2 == 0 {
					ys--
				} else {
					yr--
				}
			}

			if i%2 == 0 {
				g2.addPos(xs, ys)
			} else {
				g2.addPos(xr, yr)
			}

			g.addPos(x, y)
		}
	}

	sum1 = len(g.cells)
	sum2 = len(g2.cells)
	fmt.Println("Part1: " + fmt.Sprint(sum1))
	fmt.Println("Part2: " + fmt.Sprint(sum2))
	log.Printf("Execution time: %s", time.Since(start))
}
