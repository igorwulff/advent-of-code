package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("./sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	strength := map[string]int{"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2}
	sum := 0
	sets := make(map[string]int) // "32T3K" => 765

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(data[1])
		sets[data[0]] = value

	}
	setKeys := make([]string, 0, len(sets))

	for key := range sets {
		setKeys = append(setKeys, key)

		matches := make(map[string]int)
		for _, v := range key {
			matches[string(v)]++
		}

		// Sort map by value.
		keys := make([]string, 0, len(matches))
		for v2 := range matches {
			keys = append(keys, v2)
		}

		sort.SliceStable(keys, func(i int, j int) bool {
			if matches[keys[i]] == matches[keys[j]] {
				return strength[keys[i]] > strength[keys[j]]
			}

			return matches[keys[i]] > matches[keys[j]]
		})
		// End of sorting.

		for _, v := range keys {
			fmt.Print(v + ":")
			fmt.Println(matches[v])
		}

		fmt.Println("==============")
	}

	sort.SliceStable(keys, func(i int, j int) bool {
		if matches[keys[i]] == matches[keys[j]] {
			return strength[keys[i]] > strength[keys[j]]
		}

		return matches[keys[i]] > matches[keys[j]]
	})

	fmt.Println("The answer is: `" + fmt.Sprint(sum) + "`")

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}
