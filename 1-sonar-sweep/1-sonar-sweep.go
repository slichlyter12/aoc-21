package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	increases := 0
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	depths := strings.Split(string(input), "\n")
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
