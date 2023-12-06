package lib

import "testing"

func TestMapFindMappingWithValueFound(t *testing.T) {
	mapping := Map{
		DestinationRangeStart: 0,
		SourceRangeStart:      1,
		RangeLength:           2,
	}

	mappingResult := mapping.FindMapping(2)

	if mappingResult != 1 {
		t.Errorf("Mapping result is not the expected one. Expected: %v, actual: %v", 1, mappingResult)
	}
}

func TestMapFindMappingWithValueNotFound(t *testing.T) {
	mapping := Map{
		DestinationRangeStart: 0,
		SourceRangeStart:      1,
		RangeLength:           2,
	}

	mappingResult := mapping.FindMapping(3)

	if mappingResult != -1 {
		t.Errorf("Mapping result is not the expected one. Expected: %v, actual: %v", -1, mappingResult)
	}
}
