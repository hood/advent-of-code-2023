package shared

func FindLowestInteger(values []int) int {
	result := -1

	for _, value := range values {
		if result == -1 || value < result {
			result = value
		}
	}

	return result
}
