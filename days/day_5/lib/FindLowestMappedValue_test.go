package lib

import "testing"

func TestFindLowestMappedValue(t *testing.T) {
	mappings := []Map{
		{
			DestinationRangeStart: 50,
			SourceRangeStart:      98,
			RangeLength:           2,
		},
		{
			DestinationRangeStart: 52,
			SourceRangeStart:      50,
			RangeLength:           48,
		},
	}

	result := FindLowestMappedValue(79, mappings)

	if result != 81 {
		t.Errorf("Result is not the expected one. Expected: %v, actual: %v", 81, result)
	}
}
