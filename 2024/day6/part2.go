package main

import (
	"fmt"
	"sync"

	"github.com/igorwulff/advent-of-code/2024/day6/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	grid, guard := shared.ParseInput(input)

	type Job struct {
		X, Y int
	}

	size := shared.Width * shared.Height

	jobs := make(chan Job, size)
	results := make(chan bool, size)

	// Worker function
	worker := func(jobs <-chan Job, results chan<- bool) {
		for job := range jobs {
			copiedGrid := shared.Grid{
				ObstacleX: job.X,
				ObstacleY: job.Y,
			}
			copiedGuard := shared.Guard{
				X:       guard.X,
				Y:       guard.Y,
				Path:    make([]int, 0, 100),
				Visited: make(map[int]shared.Dir, 100),
			}

			for {
				if m, err := copiedGuard.Move(copiedGrid); !m {
					if err != nil {
						results <- true
					}
					break
				}
			}

			results <- false
		}
	}

	// Launch workers
	var wg sync.WaitGroup
	wg.Add(size)
	for range size {
		go func() {
			defer wg.Done()
			worker(jobs, results)
		}()
	}

	// Enqueue jobs
	go func() {
		for y := 0; y < shared.Height; y++ {
			for x := 0; x < shared.Width; x++ {
				// Skip starting position
				if x == guard.X && y == guard.Y {
					continue
				}

				// Skip obstacles
				if grid.IsObstacle(x, y) {
					continue
				}

				jobs <- Job{X: x, Y: y}
			}
		}
		close(jobs) // Close jobs channel when all jobs are enqueued
	}()

	// Close results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	stuck := 0
	for result := range results {
		if result {
			stuck++
		}
	}

	return fmt.Sprint(stuck)
}
