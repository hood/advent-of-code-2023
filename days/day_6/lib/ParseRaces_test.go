package lib

import "testing"

func TestParseRaces(t *testing.T) {
	races := ParseRaces([]string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	})

	expected := []Race{
		{
			Time:     7,
			Distance: 9,
		},
		{
			Time:     15,
			Distance: 40,
		},
		{
			Time:     30,
			Distance: 200,
		},
	}

	for i, expectedResult := range expected {
		if races[i].Time != expectedResult.Time ||
			races[i].Distance != expectedResult.Distance {
			t.Errorf("Expected %v, got %v", expectedResult, races[i])
		}
	}
}
