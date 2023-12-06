package lib

func FindLowestMappedValueInRange(seedsRange SeedsRange, mappings []Map) int {
	result := -1

	for i := seedsRange.Start; i <= seedsRange.End; i++ {
		mappedValue := finder(i, mappings)

		if result == -1 || mappedValue < result {
			result = mappedValue
		}
	}

	return result
}

func finder(value int, mappings []Map) int {
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

	return result
}
