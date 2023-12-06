package lib

func FindLowestMappedValue(value int, mappings []Map) int {
	result := -1

	for _, mapping := range mappings {
		mappedValue := mapping.FindMapping(value)

		if mappedValue == -1 {
			continue
		}

		if result == -1 || mappedValue < result {
			result = mappedValue
		}
	}

	if result == -1 {
		result = value
	}

	return result
}
