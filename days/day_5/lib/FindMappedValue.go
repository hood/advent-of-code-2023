package lib

func FindMappedValue(value int, mappings [][]int) int {
	result := -1
	// resultChannel := make(chan int)

	println("OK")

	for _, pair := range mappings {
		// go func(pair []int) {
		println("Pair", pair[0], pair[1])

		lowestMappedValue := pair[0]

		if lowestMappedValue > value {
			continue
		}

		if pair[0] != value {
			continue
		}

		// resultChannel <- pair[1]
		// }(pair)
	}

	// go func() {
	// 	for foundValue := range resultChannel {
	// 		if result == -1 || foundValue < result {
	// 			result = foundValue
	// 		}
	// 	}
	// }()

	// go func() {
	// 	defer close(resultChannel)
	// }()

	if result == -1 {
		result = value
	}

	return result
}
