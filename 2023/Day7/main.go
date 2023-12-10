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

var strength = map[string]int{"A": 14, "K": 13, "Q": 12, "J": 1, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2}

func main() {
	start := time.Now()

	file, err := os.Open("./sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	hands := make(map[string]int) // "32T3K" => 765
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(data[1])

		i := data[0]
		hands[i] = value
	}

	ranking := make([]string, 0, len(hands))
	for hand := range hands {
		ranking = append(ranking, hand)
	}

	sort.SliceStable(ranking, func(i int, j int) bool {
		left := make(map[rune]int)
		sortedLeft := sortHand(ranking[i])
		for _, v := range sortedLeft {
			left[v]++
		}

		right := make(map[rune]int)
		sortedRight := sortHand(ranking[j])
		for _, v := range sortedRight {
			right[v]++
		}

		valueLeft := handValue(left)
		valueRight := handValue(right)

		if valueLeft == valueRight {
			for k := 0; k < 5; k++ {
				if strength[string(ranking[i][k])] == strength[string(ranking[j][k])] {
					continue
				}

				return strength[string(ranking[i][k])] < strength[string(ranking[j][k])]
			}
		}
		return valueLeft < valueRight
	})

	for i, v := range ranking {
		fmt.Print(i)
		fmt.Print(":" + v + ":")
		fmt.Print(hands[v])
		fmt.Println()
		sum += hands[v] * (i + 1)
	}

	fmt.Println("The answer is: `" + fmt.Sprint(sum) + "`")

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func handValue(hand map[rune]int) int {

	value := 0 //XX.XX.XX.XX.XX.XX.XX

	// Get Jokers.
	jokers := 0
	for k := range hand {
		if string(k) == "J" {
			jokers++
		}
	}

	// Hand always has 5 cards.
	for k, v := range hand {
		fmt.Print(string(k) + ":")
		fmt.Println(v)
		// Five of a kind.
		if v+jokers == 5 {
			value = 7
		}

		// Four of a kind.
		if v+jokers == 4 {
			value = 6
		}

		// Three of a kind
		if v+jokers == 3 {
			if value == 2 {
				value = 5
			} else {
				value = 4
			}
		}

		if v == 2 {
			if value == 4 { // if value is already three of a kind then:
				value = 5 // Full house... = 5
			} else if value == 2 { //if value is already a pair.
				value = 3 // 2 pairs... = 3
			} else {
				value = 2
			}
		}
	}

	return value
}

// Sort a hand by value.
func sortHand(hand string) string {
	matches := make(map[string]int)
	for _, card := range hand {
		matches[string(card)]++
	}

	// Sort map by value.
	cards := make([]string, 0, len(matches))
	for card := range matches {
		cards = append(cards, card)
	}

	chars := []rune(hand)
	sort.SliceStable(chars, func(i, j int) bool {
		if matches[string(chars[i])] == matches[string(chars[j])] {
			return strength[string(chars[i])] > strength[string(chars[j])]
		}
		return matches[string(chars[i])] > matches[string(chars[j])]
	})

	return string(chars)
}
