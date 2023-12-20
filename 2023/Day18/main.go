package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int64

	plan := make(map[string]int64, 1000000)

	var x, y int64
	plan["X:"+string(x)+",Y:"+string(y)] = 0

	steps := make([][]int64, 0)
	for scanner.Scan() {
		raw := scanner.Text()
		text := strings.Split(raw, " ")
		//n, _ := strconv.Atoi(text[1])
		//hex := text[2]

		n, _ := strconv.ParseInt(text[2][2:len(text[2])-2], 16, 64)
		dir, _ := strconv.ParseInt(text[2][len(text[2])-2:len(text[2])-1], 16, 64)

		// Between value... Start Length?
		if dir == 0 {

			steps = append(steps, []int64{1, 0, n})

		} else if dir == 2 {
			x -= n
			steps = append(steps, []int64{-1, 0, n})

		} else if dir == 1 {
			y += n

			steps = append(steps, []int64{0, 1, n})
		} else if dir == 3 {
			y -= n

			steps = append(steps, []int64{0, -1, n})

		}
	}

	sum = calc(steps)

	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func calc(steps [][]int64) int64 {
	var pos int64
	ans := 1.0

	for _, step := range steps {
		x := step[0]
		y := step[1]
		n := step[2]

		pos += x * n
		ans += float64(y*n*pos) + float64(n)/2.0
	}

	return int64(ans)
}
