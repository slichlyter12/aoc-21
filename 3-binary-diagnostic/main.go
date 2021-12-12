package main

import (
	"fmt"

	"github.com/slichlyter12/aoc-21/3-binary-diagnostic/diagnostic"
	"github.com/slichlyter12/aoc-21/common"
)

func main() {
	report := common.GetFileStrs("3-binary-diagnostic/input.txt")
	gammaRate := diagnostic.GammaRate(report)
	epsilonRate := diagnostic.EpsilonRate(gammaRate)

	gammaRateDec := diagnostic.BinaryToInt(gammaRate)
	epsilonRateDec := diagnostic.BinaryToInt(epsilonRate)

	fmt.Printf("Gamma Rate:   %v (%v)\n", gammaRate, gammaRateDec)
	fmt.Printf("Epsilon Rate: %v (%v)\n", epsilonRate, epsilonRateDec)
	fmt.Println("Power Consumption: ", gammaRateDec*epsilonRateDec)
}
