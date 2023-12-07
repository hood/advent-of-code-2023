package lib

import "testing"

func TestFindLowestMappedValueInRange(t *testing.T) {
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

	result := -1

	r := SeedsRange{
		Start: 79,
		End:   13 - 1,
	}

	for i := r.Start; i <= r.End; i++ {
		mappedValue := finder(i, mappings)

		if mappedValue == -1 {
			continue
		}

		if result == -1 || mappedValue < result {
			result = mappedValue
		}
	}

	if result != 84 {
		t.Errorf("Result is not the expected one. Expected: %v, actual: %v", 84, result)
	}
}
