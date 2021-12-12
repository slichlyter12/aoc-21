package main

import (
	"fmt"

	"github.com/slichlyter12/aoc-21/2-dive/plotcourse"
	"github.com/slichlyter12/aoc-21/common"
)

func main() {
	directions := common.GetFileStrs("2-dive/input.txt")
	var instructions []plotcourse.Instruction
	for _, direction := range directions {
		instructions = append(instructions, plotcourse.GetInstruction(direction))
	}
	hPos, depth := plotcourse.PlotCourse(instructions)
	fmt.Printf("Horizontal Position: %v\nDepth: %v\n", hPos, depth)
	fmt.Printf("Position: %v\n", hPos*depth)
}
