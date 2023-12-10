package lib

import (
	"adventofcode2023/days/shared"
	"testing"
)

func TestMap(t *testing.T) {
	input := []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}

	_, startingPoint := MapFromLines(input)

	shared.AssertEqual(t, startingPoint[0], 2)
	shared.AssertEqual(t, startingPoint[1], 0)
}

func TestFindConnectingTiles(t *testing.T) {
	input := []string{
		".|.",
		"-S-",
		".|.",
	}

	m, _ := MapFromLines(input)

	connections := m.FindConnectingTiles(Coordinates{1, 1}, Coordinates{1, 1})

	shared.AssertEqual(t, 4, len(connections))
}

func TestFindFarthestTile(t *testing.T) {
	input := []string{
		".|-",
		"-S|",
		".||",
	}

	m, _ := MapFromLines(input)

	m.Print(Coordinates{1, 1})

	distance := m.FindFarthestTile(Coordinates{1, 1})

	shared.AssertEqual(t, 4, distance)
}
