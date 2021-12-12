package sonarsweep

func StandardIncreases(depths []int) int {
	increases := 0

	var lastDepth int
	for index, depth := range depths {
		if index == 0 {
			lastDepth = depth
		}

		if depth > lastDepth {
			increases++
		}

		lastDepth = depth
	}

	return increases
}

func SlidingWindowIncreases(depths []int) int {
	increases := 0

	var window1 = []int{
		depths[0],
		depths[1],
		depths[2],
	}
	var window2 = []int{
		depths[1],
		depths[2],
		depths[3],
	}

	for index, depth := range depths {
		if index < 3 {
			continue
		}

		if windowSum(window2) > windowSum(window1) {
			increases++
		}

		window1 = window2
		window2 = slideWindow(window2, depth)

	}

	return increases
}

func slideWindow(window []int, newDepth int) []int {
	newWindow := []int{window[1], window[2], 0}
	newWindow[2] = newDepth

	return newWindow
}

func windowSum(window []int) int {
	sum := 0
	for _, value := range window {
		sum += value
	}

	return sum
}
