package lib

func FindMappedValue(value int, mappings []Map) int {
	result := -1

	for _, mapping := range mappings {
		r := mapping.FindMapping(value)

		if result == -1 || r < result {
			result = r
		}
	}

	if result == -1 {
		result = value
	}

	return result
}
