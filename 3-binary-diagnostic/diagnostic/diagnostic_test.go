package diagnostic

import "testing"

func TestGammaRate(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	trueGammaRate := "10110"
	gammaRate := GammaRate(input)

	if gammaRate != trueGammaRate {
		t.Errorf("Got %v, wanted %v", gammaRate, trueGammaRate)
	}
}

func TestOxygenRating(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	trueOxygenRating := "10111"
	oxygenRating := OxygenRating(input)

	if oxygenRating != trueOxygenRating {
		t.Errorf("Got %v, wanted %v", oxygenRating, trueOxygenRating)
	}
}
