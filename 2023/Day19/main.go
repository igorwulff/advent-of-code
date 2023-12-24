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
	value    int64  // 10
	symbol   string // x
	next     string
}

func main() {
	start := time.Now()

	file, err := os.Open("./sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	r, _ := regexp.Compile("([a-z0-9]+){(.*)}")
	subr, _ := regexp.Compile("([a-z]+)(<|>)([0-9]+)[:]([a-zA-Z]+)")

	var sum int64
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
					parts[part[1]] = append(parts[part[1]], rule{operator: m[2], value: int64(v), symbol: m[1], next: m[4]})
				}
			}
		}
	}

	sum = find(parts, map[string]int64{"x": 0, "xx": 4000, "m": 0, "mm": 4000, "a": 0, "aa": 4000, "s": 0, "ss": 4000}, "in")

	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func find(parts map[string][]rule, mx map[string]int64, nextRule string) int64 {
	var sum int64
	for _, rule := range parts[nextRule] {
		tmx := map[string]int64{"x": mx["x"], "xx": mx["xx"], "m": mx["m"], "mm": mx["mm"], "a": mx["a"], "aa": mx["aa"], "s": mx["s"], "ss": mx["ss"]}

		if rule.next == "R" {
			if rule.operator == ">" {
				tmx[rule.symbol+rule.symbol] = rule.value
			} else if rule.operator == "<" {
				tmx[rule.symbol] = rule.value - 1
			}
		} else {
			if rule.operator == ">" {
				tmx[rule.symbol] = rule.value
			} else if rule.operator == "<" {
				tmx[rule.symbol+rule.symbol] = rule.value - 1
			}
		}

		if rule.next == "R" {
			continue
		}

		if rule.next == "A" {
			sum += (tmx["xx"] - tmx["x"]) * (tmx["mm"] - tmx["m"]) * (tmx["aa"] - tmx["a"]) * (tmx["ss"] - tmx["s"])
		} else {
			sum += find(parts, tmx, rule.next)
		}
	}

	return sum
}
