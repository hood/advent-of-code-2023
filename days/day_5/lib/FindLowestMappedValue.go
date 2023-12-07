package lib

func FindLowestMappedValue(value int, mappings []Map) int {
	low, high := 0, len(mappings)-1
	result := -1

	for low <= high {
		mid := low + (high-low)/2

		if mappings[mid].SourceRangeStart <= value {
			result = mappings[mid].FindMapping(value)
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if result == -1 {
		return value
	}

	return result
}
