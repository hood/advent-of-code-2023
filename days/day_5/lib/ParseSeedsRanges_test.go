package lib

import "testing"

func TestParseSeedsRanges(t *testing.T) {
	input := []string{
		"seeds: 79 14 55 13",
		"",
	}

	result := ParseSeedsRanges(input)

	expected := []SeedsRange{
		SeedsRange{Start: 79, End: 79 + 14},
		SeedsRange{Start: 55, End: 55 + 13},
	}

	if len(result) != len(expected) {
		t.Errorf("Expected %v ranges, got %v ranges", expected, result)
	}

	for i := 0; i < len(result); i++ {
		if result[i].Start != expected[i].Start {
			t.Errorf("Expected range #%v to have start %v, got %v", i, expected[i].Start, result[i].Start)
		}

		if result[i].End != expected[i].End {
			t.Errorf("Expected range #%v to have end %v, got %v", i, expected[i].End, result[i].End)
		}
	}
}
