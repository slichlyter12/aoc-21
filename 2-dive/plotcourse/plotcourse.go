package plotcourse

import (
	"strings"

	"github.com/slichlyter12/aoc-21/common"
)

type Instruction struct {
	Direction string
	Magnitude int
}

// PlotCourse calculates a horizontal position and depth given
// a set of instructions
func PlotCourse(instructions []Instruction) (hPos int, depth int) {
	hPos = 0
	depth = 0

	for _, instruction := range instructions {
		switch instruction.Direction {
		case "forward":
			hPos += instruction.Magnitude
		case "up":
			depth -= instruction.Magnitude
		case "down":
			depth += instruction.Magnitude
		}
	}

	return hPos, depth
}

// PlotCourseWithAim calculates horizontal position and depth given
// a set of instructions but uses aim to get there
func PlotCourseWithAim(instructions []Instruction) (hPos int, depth int, aim int) {
	hPos = 0
	depth = 0
	aim = 0

	for _, instruction := range instructions {
		switch instruction.Direction {
		case "forward":
			hPos += instruction.Magnitude
			depth += aim * instruction.Magnitude
		case "up":
			aim -= instruction.Magnitude
		case "down":
			aim += instruction.Magnitude
		}
	}

	return hPos, depth, aim
}

// GetInstruction converts a string of instructions into an Instruction object
func GetInstruction(instructionStr string) Instruction {
	instructionSlice := strings.Fields(instructionStr)
	instruction := Instruction{
		instructionSlice[0],
		common.Atoi(instructionSlice[1]),
	}

	return instruction
}
