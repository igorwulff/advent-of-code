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

func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	list := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		list = append(list, strings.Split(text, ",")...)
	}

	boxes := make([]Box, 256)

	r, _ := regexp.Compile("^([a-z]+)([=|-]+)([0-9]*)")
	// rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
	for _, sequence := range list {

		text := r.FindStringSubmatch(sequence)

		hash := getHash(text[1])
		label := text[1]

		operator := text[2]

		if operator == "=" {
			value, _ := strconv.Atoi(text[3])
			lens := Lens{orig: text[0], label: label, value: value}
			isReplaced := false
			for _, oldLens := range boxes[hash].list {
				if oldLens.label == label {
					oldLens.value = value
					isReplaced = true
				}
			}

			if !isReplaced {
				boxes[hash].list = append(boxes[hash].list, &lens)
			}
		} else if operator == "-" {
			found := false
			i := 0
			for _, oldLens := range boxes[hash].list {
				if oldLens.label == label {
					found = true
					break
				}

				i++
			}

			if found {
				boxes[hash].list = remove(boxes[hash].list, i)
			}
		}
	}

	for i, box := range boxes {
		for j, lens := range box.list {
			sum += (i + 1) * (j + 1) * lens.value
		}
	}

	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func remove(slice []*Lens, s int) []*Lens {
	return append(slice[:s], slice[s+1:]...)
}

type Box struct {
	list []*Lens
}

type Lens struct {
	orig  string
	label string
	value int
}

func getHash(label string) int {
	v := 0
	for _, r := range label {
		v += int(r)
		v *= 17
		v %= 256
	}
	return v
}
