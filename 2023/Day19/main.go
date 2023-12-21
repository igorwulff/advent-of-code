package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// x>10:one
type rule struct {
	operator string // >
	value    int    // 10
	symbol   string // x
	next     string
}

func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	r, _ := regexp.Compile("([a-z]+){(.*)}")
	subr, _ := regexp.Compile("([a-z]+)(<|>)([0-9]+)[:]([a-zA-Z]+)")
	ratintr, _ := regexp.Compile("([0-9]+)")

	sum := 0
	init := true
	parts := make(map[string][]rule, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			init = false
			continue
		}

		if init {
			part := r.FindStringSubmatch(line)
			for _, text := range strings.Split(part[2], ",") {
				m := subr.FindStringSubmatch(text)

				if len(m) == 0 {
					parts[part[1]] = append(parts[part[1]], rule{next: text})
				} else {
					v, _ := strconv.Atoi(m[3])
					parts[part[1]] = append(parts[part[1]], rule{operator: m[2], value: v, symbol: m[1], next: m[4]})
				}
			}
		} else {
			ratingsUnparsed := ratintr.FindAllString(line, -1)
			ratings := make(map[string]int, 0)

			ratings["x"], _ = strconv.Atoi(ratingsUnparsed[0])
			ratings["m"], _ = strconv.Atoi(ratingsUnparsed[1])
			ratings["a"], _ = strconv.Atoi(ratingsUnparsed[2])
			ratings["s"], _ = strconv.Atoi(ratingsUnparsed[3])
			fmt.Println(ratings)
			if find(ratings, parts, "in", 0) {
				value := ratings["x"] + ratings["m"] + ratings["a"] + ratings["s"]
				fmt.Println(value)
				sum += value
			}
		}
	}

	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func find(ratings map[string]int, parts map[string][]rule, nextRule string, step int) bool {
	rules := parts[nextRule]
	for _, rule := range rules {
		if rule.operator == "" {
			if rule.next == "A" {
				return true
			} else if rule.next == "R" {
				return false
			}
			return find(ratings, parts, rule.next, step+1)
		} else {
			if rule.operator == ">" {
				if ratings[rule.symbol] > rule.value {
					if rule.next == "A" {
						return true
					} else if rule.next == "R" {
						return false
					}
					return find(ratings, parts, rule.next, step+1)
				}
			} else {
				if ratings[rule.symbol] < rule.value {
					if rule.next == "A" {
						return true
					} else if rule.next == "R" {
						return false
					}
					return find(ratings, parts, rule.next, step+1)
				}
			}
		}
	}

	return false
}
