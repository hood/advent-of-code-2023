package lib

func FindLowestMappedValueInRange(seedsRange SeedsRange, mappings []Map) int {
	result := -1

	for i := seedsRange.Start; i <= seedsRange.End; i++ {
		mappedValue := FindLowestMappedValue(i, mappings)

		if result == -1 || mappedValue < result {
			result = mappedValue
		}
	}

	return result
}
