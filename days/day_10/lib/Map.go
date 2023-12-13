package lib

import (
	"fmt"
)

type Map [][]rune

type Coordinates struct {
	X int
	Y int
}

func (c *Coordinates) SameAs(oc Coordinates) bool {
	if c.X == oc.X && c.Y == oc.Y {
		return true
	}

	return false
}

func MapFromLines(lines []string) (Map, Coordinates) {
	m := Map{}
	startingPoint := Coordinates{}

	for rowindex, line := range lines {
		row := []rune{}

		for columnindex, char := range line {
			if char == 'S' {
				startingPoint = Coordinates{
					X: columnindex,
					Y: rowindex,
				}
			}

			row = append(row, char)
		}

		m = append(m, row)
	}

	return m, startingPoint
}

func (m *Map) FindFarthestTile(startingPoint Coordinates) int {
	currentPosition := Coordinates{Y: startingPoint.Y, X: startingPoint.X}
	previousPosition := Coordinates{Y: startingPoint.Y, X: startingPoint.X}

	distance := 0

	// Loop around the current cell.
	// Find all pipes that connect, excluding the previous one.
	for {
		println("")
		m.Print(currentPosition, previousPosition)

		// Find the connecting tiles.
		connecting := m.FindConnectingTiles(
			currentPosition,
			previousPosition,
		)

		fmt.Printf("Connecting tiles: %v\n", len(connecting))
		for _, connection := range connecting {
			println("Connection(x,y):", connection.X, connection.Y)
		}
		println("Previous(x,y):", previousPosition.X, previousPosition.Y)

		// If there are no connecting tiles, we're done.
		if len(connecting) == 0 {
			break
		}

		if distance > 1 && len(connecting) > 1 {
			panic("We can't handle this!")
		}

		// For every connecting tile we find (there should be max 1), move to that
		// tile.
		previousPosition.X = currentPosition.X
		previousPosition.Y = currentPosition.Y
		for _, connection := range connecting {
			currentPosition.X = connection.X
			currentPosition.Y = connection.Y
		}

		distance++

		if distance > len(*m)*len((*m)[0]) {
			panic("This went too far!")
		}
	}

	return distance
}

// Given a point, and another point that we came from, find all tiles that
// are adjacent to the point, connect to it, and are not the previous point.
func (m *Map) FindConnectingTiles(
	point Coordinates,
	previous Coordinates,
) []Coordinates {
	connections := []Coordinates{}

	// Check all directions of the point's endpoints.
	for _, direction := range Tiles[m.At(point)] {
		// West boundary.
		if point.Y == 0 && direction.Y == -1 {
			continue
		}

		// South boundary.
		if point.Y == len(*m)-1 && direction.Y == 1 {
			continue
		}

		// North boundary.
		if point.X == 0 && direction.X == -1 {
			continue
		}

		// East boundary.
		if point.X == len((*m)[0])-1 && direction.X == 1 {
			continue
		}

		// Calculate the position of the tile in the current direction.
		position := Coordinates{
			X: point.X + direction.X,
			Y: point.Y + direction.Y,
		}

		// Ignore the previous point.
		if position.SameAs(previous) {
			continue
		}

		// Get the symbol for the calculated position.
		current := m.At(position)

		// Ignore spaces.
		if current == '.' {
			continue
		}

		// Loop the endpoints of the current tile, check if they connect to the
		// original tile.
		for _, endpoint := range Tiles[current] {
			if endpoint.X == -direction.X &&
				endpoint.Y == -direction.Y {
				connections = append(connections, position)

				// As soon as one connects, we're done checking.
				break
			}
		}
	}

	return connections
}

func (m *Map) At(coordinates Coordinates) rune {
	return (*m)[coordinates.Y][coordinates.X]
}

func (m *Map) Print(highlight Coordinates, previousHighlight Coordinates) {
	for verticalIndex, row := range *m {
		for horizontalIndex, tile := range row {
			output := string(tile)

			if verticalIndex == highlight.Y && horizontalIndex == highlight.X {
				printRed(output)
			} else if verticalIndex == previousHighlight.Y && horizontalIndex == previousHighlight.X {
				printYellow(output)
			} else {
				fmt.Print(output)
			}
		}

		println("")
	}
}

func printRed(v any) {
	const colorRed = "\033[0;31m"
	const colorNone = "\033[0m"

	fmt.Printf("%v%v%v", colorRed, v, colorNone)
}

func printYellow(v any) {
	const colorYellow = "\033[0;33m"
	const colorNone = "\033[0m"

	fmt.Printf("%v%v%v", colorYellow, v, colorNone)
}
