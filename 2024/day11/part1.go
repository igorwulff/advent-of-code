package main

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/igorwulff/advent-of-code/2024/day11/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	values := strings.Split(input, " ")

	var sum int64 = 0
	memo := make(map[string]int)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for _, value := range values {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&sum, int64(shared.Blink(25, value, &memo, &mutex)))
		}()
	}
	wg.Wait()

	return fmt.Sprint(sum)
}
