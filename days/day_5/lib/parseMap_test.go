package lib

import "testing"

func TestParseMap(t *testing.T) {
	mapString := "0 1 2"
	expectedMap := Map{
		DestinationRangeStart: 0,
		SourceRangeStart:      1,
		RangeLength:           2,
	}

	mapParsed := ParseMap(mapString)

	if mapParsed != expectedMap {
		t.Errorf("Map parsed is not the expected one. Expected: %v, actual: %v", expectedMap, mapParsed)
	}
}
