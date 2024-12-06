package main

import (
	"fmt"
	"sync"

	"github.com/igorwulff/advent-of-code/2024/day6/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	grid, guard := shared.ParseInput(input)

	startx := guard.X
	starty := guard.Y

	type Job struct {
		X, Y int
	}

	jobs := make(chan Job, shared.Width*shared.Height)
	results := make(chan bool, shared.Width*shared.Height)

	// Worker function
	worker := func(jobs <-chan Job, results chan<- bool) {
		for job := range jobs {
			// Skip the starting position
			if job.X == startx && job.Y == starty {
				results <- false
				continue
			}

			// Skip obstacles
			if grid.IsObstacle(job.X, job.Y) {
				results <- false
				continue
			}

			copiedGrid := shared.Grid{
				ObstacleX: job.X,
				ObstacleY: job.Y,
			}
			copiedGuard := shared.Guard{
				X:       startx,
				Y:       starty,
				Path:    make([]int, 0),
				Visited: make(map[int]shared.Dir),
			}

			// Set obstacle and simulate guard movement
			copiedGrid.ObstacleX = job.X
			copiedGrid.ObstacleY = job.Y

			stuck := false
			for {
				moved, err := copiedGuard.Move(copiedGrid)
				if !moved {
					if err != nil {
						stuck = true
					}
					break
				}
			}

			results <- stuck
		}
	}

	// Launch workers
	numWorkers := shared.Width * shared.Height
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, results)
		}()
	}

	// Enqueue jobs
	go func() {
		for y := 0; y < shared.Height; y++ {
			for x := 0; x < shared.Width; x++ {
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
	stuckCount := 0
	for result := range results {
		if result {
			stuckCount++
		}
	}

	return fmt.Sprint(stuckCount)
}
