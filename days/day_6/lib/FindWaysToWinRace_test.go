package lib

import "testing"

func TestFindWaysToWinRace(t *testing.T) {
	race := Race{
		Time:     7,
		Distance: 9,
	}

	expected := []int{2, 3, 4, 5}

	result := FindWaysToWinRace(&race)

	if len(result) != len(expected) {
		t.Errorf("Expected %v results, got %v results", expected, result)
	}

	// Since we expect the results to be in order...
	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], result[i])
		}
	}
}
