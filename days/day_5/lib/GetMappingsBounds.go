package lib

func GetMappingsDestinationsBounds(mappings []Map) (int, int) {
	lower := mappings[0].DestinationRangeStart
	upper := mappings[len(mappings)-1].DestinationRangeStart + mappings[len(mappings)-1].RangeLength

	return lower, upper
}
