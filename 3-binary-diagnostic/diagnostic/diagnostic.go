package diagnostic

func GammaRate(report []string) string {
	numBits := len(report[0])
	var gammaRate []rune

	for bitPos := 0; bitPos < numBits; bitPos++ {
		mostCommon := commonBit(report, bitPos, '0', true)
		gammaRate = append(gammaRate, mostCommon)
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

func OxygenRating(report []string) string {
	slimmedReport := report
	bitPos := 0
	for len(slimmedReport) != 1 {
		mostCommonBit := commonBit(slimmedReport, bitPos, '1', true)
		slimmedReport = keepDesiredBins(slimmedReport, bitPos, mostCommonBit)
		bitPos++
	}

	return slimmedReport[0]
}

func CO2Rating(report []string) string {
	slimmedReport := report
	bitPos := 0

	for len(slimmedReport) != 1 {
		leastCommonBit := commonBit(slimmedReport, bitPos, '0', false)
		slimmedReport = keepDesiredBins(slimmedReport, bitPos, leastCommonBit)
		bitPos++
	}

	return slimmedReport[0]
}

func commonBit(report []string, bitIndex int, preferedBit rune, mostCommon bool) rune {
	numZeros := 0
	numOnes := 0

	for binIndex := 0; binIndex < len(report); binIndex++ {
		if rune(report[binIndex][bitIndex]) == '0' {
			numZeros++
		} else {
			numOnes++
		}
	}

	if numOnes == numZeros {
		return preferedBit
	}

	if !mostCommon {
		if numOnes < numZeros {
			return '1'
		}

		return '0'
	}

	if numOnes > numZeros {
		return '1'
	}

	return '0'
}

func keepDesiredBins(report []string, bitIndex int, desiredBit rune) []string {
	var desiredBins []string
	for binIndex := 0; binIndex < len(report); binIndex++ {
		if rune(report[binIndex][bitIndex]) == desiredBit {
			desiredBins = append(desiredBins, report[binIndex])
		}
	}

	return desiredBins
}
