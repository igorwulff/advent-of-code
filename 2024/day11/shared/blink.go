package shared

import (
	"fmt"
	"strconv"
	"sync"
)

func Blink(i int, value string, memo *map[string]int, mutex *sync.Mutex) int {
	if i == 0 {
		return 1
	}

	i--

	if value == "0" {
		return Blink(i, "1", memo, mutex)
	}

	if len(value)%2 == 0 {
		// Generate a unique key for the current state
		key := fmt.Sprintf("%d|%s", i, value)

		// Check if the result is already computed
		mutex.Lock()
		if result, exists := (*memo)[key]; exists {
			mutex.Unlock()
			return result
		}
		mutex.Unlock()

		mid := len(value) / 2
		r, _ := strconv.Atoi(value[mid:])
		result := Blink(i, value[:mid], memo, mutex) + Blink(i, strconv.Itoa(r), memo, mutex)

		mutex.Lock()
		(*memo)[key] = result
		mutex.Unlock()

		return result
	}

	number, _ := strconv.Atoi(value)
	return Blink(i, strconv.Itoa(number*2024), memo, mutex)
}
