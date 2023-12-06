package lib

func FindMappedValue(value int, mappings [][]int) int {
	result := -1

	for _, pair := range mappings {
		lowestMappedValue := pair[0]

		if lowestMappedValue > value {
			break
		}

		if pair[0] != value {
			continue
		}

		result = pair[1]
	}

	if result == -1 {
		result = value
	}

	return result
}
