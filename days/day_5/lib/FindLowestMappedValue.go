package lib

// FindLowestMappedValue finds, given an input value, the lowest mapped value
// in a list of mappings, using binary search. If the input value is not
// mapped, it returns the input value.
func FindLowestMappedValue(value int, mappings []Map) int {
	// If value is under the lowest mapped value, it's not mapped.
	if value < mappings[0].SourceRangeStart {
		return value
	}

	// If value is over the highest mapped value, it's not mapped.
	if value > mappings[len(mappings)-1].SourceRangeStart+mappings[len(mappings)-1].RangeLength {
		return value
	}

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
