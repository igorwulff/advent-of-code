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

var strength = map[string]int{"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2}

func main() {
	start := time.Now()

	file, err := os.Open("./sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	sets := make(map[string]int)   // "32T3K" => 765
	values := make(map[string]int) // "32T3K" => 765
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(data[1])
		sets[data[0]] = value
		values[data[0]] = 0
	}

	keys := make([]string, 0, len(sets))
	for v2 := range sets {
		keys = append(keys, v2)
	}
	fmt.Println(keys)
	sort.SliceStable(keys, func(i int, j int) bool {
		sortl, valuesl := sortHand(keys[i], keys, sets)
		sortr, valuesr := sortHand(keys[j], keys, sets)

		fmt.Println("CHECK:")
		fmt.Println(sortl)
		fmt.Println(valuesl)
		fmt.Println(sortr)
		fmt.Println(valuesr)

		for _, i := range sortl {
			for _, j := range sortr {
				fmt.Println(valuesl[i])
				fmt.Println(valuesr[j])

			}
		}
		return true
	})

	fmt.Println("By name:", keys)
	fmt.Println("The answer is: `" + fmt.Sprint(sum) + "`")

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func sortHand(hand string, keys []string, sets map[string]int) ([]string, map[string]int) {
	matches := make(map[string]int)
	for _, v := range hand {
		matches[string(v)]++
	}

	// Sort map by value.
	cards := make([]string, 0, len(matches))
	for v := range matches {
		cards = append(cards, v)
	}

	sort.SliceStable(keys, func(i int, j int) bool {
		if matches[keys[i]] == matches[keys[j]] {
			return strength[keys[i]] < strength[keys[j]]
		}
		return matches[keys[i]] < matches[keys[j]]
	})

	return cards, matches
}
