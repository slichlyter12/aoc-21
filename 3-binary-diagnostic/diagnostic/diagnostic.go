package diagnostic

import (
	"log"
	"strconv"
)

func GammaRate(report []string) string {
	numBits := len(report[0])
	var gammaRate []rune

	for bitPos := 0; bitPos < numBits; bitPos++ {
		numZeros := 0
		numOnes := 0
		for binIndex := 0; binIndex < len(report); binIndex++ {
			bit := rune(report[binIndex][bitPos])
			if bit == '0' {
				numZeros++
			} else {
				numOnes++
			}
		}

		if numZeros > numOnes {
			gammaRate = append(gammaRate, '0')
		} else {
			gammaRate = append(gammaRate, '1')
		}
	}

	return string(gammaRate)
}

func EpsilonRate(gammaRate string) string {
	var epsilonRate []rune
	for _, bit := range gammaRate {
		if bit == '0' {
			epsilonRate = append(epsilonRate, '1')
		} else {
			epsilonRate = append(epsilonRate, '0')
		}
	}

	return string(epsilonRate)
}

func BinaryToInt(binary string) int {
	output, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatalf("Could not convert binary to integer: %v", err)
	}

	return int(output)
}
