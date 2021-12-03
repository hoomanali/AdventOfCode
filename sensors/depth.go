package sensors

import (
	"fmt"
	"strconv"

	"alihooman.com/advent2021/util"
)

type SensorWidth int

const (
	Singles SensorWidth = iota
	Triples
)

// DpethIncreases counts how many times the depth increases over the given
// input. Single width compares 1 depth at a time, triple width compares the
// sum of 3 depths at a time.
func DepthIncreases(width SensorWidth) (increaseCount int) {
	stringDepths := util.ReadInputs()
	depths := depthsToInt(stringDepths)

	if width == Singles {
		return countSingleWidth(depths)
	} else if width == Triples {
		return countTripleWidth(depths)
	} else {
		util.Handle(fmt.Errorf("incorrect depth width"))
		return 0
	}
}

func countSingleWidth(depths []int) (increaseCount int) {
	increaseCount = 0

	for idx := 1; idx < len(depths); idx++ {
		if depths[idx] > depths[idx-1] {
			increaseCount++
		}
	}
	fmt.Println(increaseCount)
	return
}

func countTripleWidth(depths []int) (increaseCount int) {
	increaseCount = 0

	for idx := 3; idx < len(depths); idx++ {
		currTrip := depths[idx] + depths[idx-1] + depths[idx-2]
		prevTrip := depths[idx-1] + depths[idx-2] + depths[idx-3]
		if currTrip > prevTrip {
			increaseCount++
		}
	}
	fmt.Println(increaseCount)
	return
}

func depthsToInt(stringDepths []string) (depths []int) {
	depths = make([]int, len(stringDepths))
	for idx, depth := range stringDepths {
		d, err := strconv.Atoi(depth)
		depths[idx] = d
		util.Handle(err)
	}
	return
}
