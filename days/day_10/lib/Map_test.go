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

	shared.AssertEqual(t, startingPoint.X, 0)
	shared.AssertEqual(t, startingPoint.Y, 2)
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

func TestFindConnectingTiles2(t *testing.T) {
	input := []string{
		"...",
		"-S-",
		"...",
	}

	m, _ := MapFromLines(input)

	connections := m.FindConnectingTiles(
		Coordinates{X: 1, Y: 1},
		Coordinates{X: 1, Y: 1},
	)

	leftConnection := Coordinates{X: 0, Y: 1}
	rightConnection := Coordinates{X: 2, Y: 1}

	shared.AssertEqual(t, 2, len(connections))

	shared.AssertEqual(t, true, connections[0].SameAs(leftConnection))
	shared.AssertEqual(t, true, connections[1].SameAs(rightConnection))
}

func TestFindConnectingTiles3(t *testing.T) {
	input := []string{
		".F.",
		".S.",
		"...",
	}

	m, _ := MapFromLines(input)

	connections := m.FindConnectingTiles(Coordinates{1, 1}, Coordinates{1, 1})

	topConnection := Coordinates{X: 1, Y: 0}

	shared.AssertEqual(t, 1, len(connections))

	shared.AssertEqual(t, true, connections[0].SameAs(topConnection))
}

func TestFindConnectingTiles4(t *testing.T) {
	input := []string{
		"...",
		"-S.",
		"...",
	}

	m, _ := MapFromLines(input)

	connections := m.FindConnectingTiles(
		Coordinates{X: 0, Y: 1},
		Coordinates{X: 1, Y: 1},
	)

	shared.AssertEqual(t, 0, len(connections))
}

// func TestFindFarthestTile(t *testing.T) {
// 	input := []string{
// 		".F-",
// 		"-S|",
// 		".||",
// 	}

// 	m, _ := MapFromLines(input)

// 	distance := m.FindFarthestTile(Coordinates{1, 1})

// 	shared.AssertEqual(t, 4, distance)
// }
