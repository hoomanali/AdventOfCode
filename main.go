package main

import (
	"fmt"
	"os"

	"alihooman.com/advent2021/sensors"
	"alihooman.com/advent2021/util"
)

func main() {
	util.FileName = os.Args[1]

	if os.Args[2] == "1" {
		fmt.Printf("\nSingle depth increases: %v\n",
			sensors.DepthIncreases(sensors.Singles))
	} else {
		fmt.Printf("\nTriple depth increases: %v\n",
			sensors.DepthIncreases(sensors.Triples))
	}
}
