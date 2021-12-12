package plotcourse

import (
	"strings"

	"github.com/slichlyter12/aoc-21/common"
)

type Instruction struct {
	Direction string
	Magnitude int
}

func Plotcourse(instructions []Instruction) (hPos int, depth int) {
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

// GetInstruction converts a string of instructions into an Instruction object
func GetInstruction(instructionStr string) Instruction {
	instructionSlice := strings.Fields(instructionStr)
	instruction := Instruction{
		instructionSlice[0],
		common.Atoi(instructionSlice[1]),
	}

	return instruction
}
