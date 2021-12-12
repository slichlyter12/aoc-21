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

	gammaRateDec := common.Btoi(gammaRate)
	epsilonRateDec := common.Btoi(epsilonRate)

	fmt.Printf("Gamma Rate:   %v (%v)\n", gammaRate, gammaRateDec)
	fmt.Printf("Epsilon Rate: %v (%v)\n", epsilonRate, epsilonRateDec)
	fmt.Println("Power Consumption: ", gammaRateDec*epsilonRateDec)

	fmt.Println("--------------")

	oxygenRating := diagnostic.OxygenRating(report)
	co2Rating := diagnostic.CO2Rating(report)

	oxygenRatingDec := common.Btoi(oxygenRating)
	co2RatingDec := common.Btoi(co2Rating)
	fmt.Printf("Oxygen Rating: %v (%v)\n", oxygenRating, oxygenRatingDec)
	fmt.Printf("Oxygen Rating: %v (%v)\n", co2Rating, co2RatingDec)
	fmt.Println("Life Support Rating: ", oxygenRatingDec*co2RatingDec)
}
