package day_10

import (
	"adventofcode2023/days/day_10/lib"
	"adventofcode2023/days/shared"
	"testing"
)

func TestDay10Part21Test(t *testing.T) {
	lines := shared.ReadFile("./part_2_test_input.txt")

	m, _ := lib.MapFromLines(lines)

	// Flood fill starting from all the cells at the edges of the map.
	for y, column := range m {
		for x, v := range m[y] {
			c := lib.Coordinates{X: x, Y: y}

			if (x == 0 || y == 0 || x == len(m[y])-1 || y == len(column)-1) &&
				(v == '.' || v == '-' || v == '7' || v == 'F') {
				lib.FloodFill(&m, c, 'X')
			}
		}
	}

	result := 0

	for y := range m {
		for _, v := range m[y] {
			if v == '.' {
				result++
			}
		}
	}

	shared.AssertEqual(t, 8, result)
}
