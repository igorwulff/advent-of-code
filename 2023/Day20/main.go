package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type module struct {
	variant string // & = Conjecture, %FlipFlop
	tpulse  bool
}

func main() {
	start := time.Now()

	file, err := os.Open("./sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r, _ := regexp.Compile("([a-zA-Z_0-9%&]+) -> (.*)")

	scanner := bufio.NewScanner(file)
	sum := 0

	modules := make(map[string][]string, 0)
	for scanner.Scan() {
		part := r.FindStringSubmatch(scanner.Text())
		key := part[1]
		config := strings.Split(part[2], ", ")

		modules[key] = append(modules[key], config...)
		fmt.Println(modules)
	}

	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}
