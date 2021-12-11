package main

import (
	"fmt"
	"strconv"

	"github.com/slichlyter12/aoc-21/common"
)

func main() {
	increases := 0
	depths := common.GetFileInputs("1-sonar-sweep/input.txt")

	var lastDepth int
	for index, depthStr := range depths {
		depth, _ := strconv.Atoi(depthStr)
		if index == 0 {
			lastDepth = depth
		}

		if depth > lastDepth {
			increases++
		}

		lastDepth = depth
	}

	fmt.Println(increases)
}
