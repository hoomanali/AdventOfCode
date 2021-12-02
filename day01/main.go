package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatalf("%v", e)
	}
}

func countIncreases() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	var depths []int
	increaseCount := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		check(err)
		depths = append(depths, depth)
	}

	for idx := 1; idx < len(depths); idx++ {
		if depths[idx] > depths[idx-1] {
			increaseCount++
		}
	}

	fmt.Printf("\nIncreases: %d\n", increaseCount)
}

func main() {
	countIncreases()
}
