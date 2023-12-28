package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum1, sum2 int

	for scanner.Scan() {
		c1 := make(chan int)
		c2 := make(chan int)

		go getHashWithZeroes(scanner.Text(), 5, c1)
		go getHashWithZeroes(scanner.Text(), 6, c2)

		sum1 = <-c1
		sum2 = <-c2
	}

	fmt.Println("Part1: " + fmt.Sprint(sum1))
	fmt.Println("Part2: " + fmt.Sprint(sum2))
	log.Printf("Execution time: %s", time.Since(start))
}

func getHashWithZeroes(text string, zerofillsize int, c chan int) {
	compare := ""
	for z := 0; z < zerofillsize; z++ {
		compare += "0"
	}

	i := 0
	for {
		i++
		hash := md5.Sum([]byte(text + fmt.Sprint(i)))
		enc := hex.EncodeToString(hash[:])
		if enc[0:zerofillsize] == compare {
			break
		}
	}

	c <- i
}
