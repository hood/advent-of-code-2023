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

// 	distance := m.FindFarthestTile(Coordinates{
// 		X: 1,
// 		Y: 1,
// 	})

// 	shared.AssertEqual(t, 4, distance)
// }

func TestFindFarthestTile2(t *testing.T) {
	input := []string{
		"...",
		"FS.",
		"LJ.",
	}

	m, _ := MapFromLines(input)

	distance := m.FindFarthestTile(Coordinates{
		X: 1,
		Y: 1,
	})

	shared.AssertEqual(t, 2, distance)
}

func TestAmato(t *testing.T) {
	input := []string{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	}

	m, _ := MapFromLines(input)

	result := m.Amato()

	shared.AssertEqual(t, 4, len(result))
}

func TestAmato2(t *testing.T) {
	input := []string{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	}

	m, _ := MapFromLines(input)

	result := m.Amato()

	shared.AssertEqual(t, 8, len(result))
}

func TestAmato3(t *testing.T) {
	input := []string{
		".F----7F7F7F7F-7....",
		".|F--7||||||||FJ....",
		".||.FJ||||||||L7....",
		"FJL7L7LJLJ||LJ.L-7..",
		"L--J.L7...LJS7F-7L7.",
		"....F-J..F7FJ|L7L7L7",
		"....L7.F7||L7|.L7L7|",
		".....|FJLJ|FJ|F7|.LJ",
		"....FJL-7.||.||||...",
		"....L---J.LJ.LJLJ...",
	}

	m, _ := MapFromLines(input)

	result := m.Amato()

	shared.AssertEqual(t, 8, len(result))
}
