package lib

import "testing"

func TestParseMaps(t *testing.T) {
	mapsStrings := []string{"0 1 2", "2 1 0"}

	parsedMaps := ParseMaps(mapsStrings)

	if len(parsedMaps) != len(mapsStrings) {
		t.Errorf("Parsed maps length is not the expected one. Expected: %v, actual: %v", len(mapsStrings), len(parsedMaps))
	}
}
