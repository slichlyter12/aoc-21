package main

import (
	"fmt"

	"github.com/slichlyter12/aoc-21/1-sonar-sweep/sonarsweep"
	"github.com/slichlyter12/aoc-21/common"
)

func main() {
	depths := common.GetFileInts("1-sonar-sweep/input.txt")
	standardDepths := sonarsweep.StandardIncreases(depths)
	slidingWindowDepths := sonarsweep.SlidingWindowIncreases(depths)

	fmt.Printf("Standard: %v\n", standardDepths)
	fmt.Printf("Sliding Window: %v\n", slidingWindowDepths)
}
