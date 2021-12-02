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

func countIncreasesSingles() {
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

func countIncreasesTriples() {
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

	// 0, 1, 2, 3, 4, 5
	// A, A, A
	//    B, B, B
	for idx := 3; idx < len(depths); idx++ {
		currTrip := depths[idx] + depths[idx-1] + depths[idx-2]
		prevTrip := depths[idx-1] + depths[idx-2] + depths[idx-3]
		if currTrip > prevTrip {
			increaseCount++
		}
	}

	fmt.Printf("\nIncreases: %d\n", increaseCount)
}

func main() {
	if os.Args[1] == "1" {
		countIncreasesSingles()
	} else {
		countIncreasesTriples()
	}
}
