package lib

import (
	"strings"
	"testing"
)

func TestGetMapsGroup(t *testing.T) {
	input := `
soil-to-fertilizer map:
0 15 37
37 52 2

fertilizer-to-water map:
49 53 8
0 11 42
42 0 2
57 2 4
`

	group := "fertilizer-to-water"

	result := GetMapsGroup(strings.Split(input, "\n"), group)

	if len(result) != 4 {
		t.Errorf("Expected 4 maps, got %v", len(result))
	}
}
