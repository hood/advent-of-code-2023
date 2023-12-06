package lib

import "testing"

func TestMapFindMappingWithValueFound(t *testing.T) {
	mapping := Map{
		DestinationRangeStart: 52,
		SourceRangeStart:      50,
		RangeLength:           48,
	}

	mappingResult := mapping.FindMapping(79)

	if mappingResult != 81 {
		t.Errorf("Mapping result is not the expected one. Expected: %v, actual: %v", 81, mappingResult)
	}
}

func TestMapFindMappingWithValueNotFound(t *testing.T) {
	mapping := Map{
		DestinationRangeStart: 52,
		SourceRangeStart:      50,
		RangeLength:           48,
	}

	mappingResult := mapping.FindMapping(14)

	if mappingResult != 14 {
		t.Errorf("Mapping result is not the expected one. Expected: %v, actual: %v", 14, mappingResult)
	}
}
