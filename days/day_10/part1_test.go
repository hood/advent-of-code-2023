package day_10

import (
	"adventofcode2023/days/day_10/lib"
	"adventofcode2023/days/shared"
	"testing"
)

func TestDay10Part1Test(t *testing.T) {
	lines := shared.ReadFile("./test_input.txt")

	m, startCoordinates := lib.MapFromLines(lines)

	distance := m.FindFarthestTile(startCoordinates)

	shared.AssertEqual(t, 4, distance)
}
