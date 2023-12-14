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
		// Find the connecting tiles.
		connecting := m.FindConnectingTiles(
			currentPosition,
			previousPosition,
		)

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

		if (distance > 1) && m.At(currentPosition) == 'S' {
			if (distance % 2) == 0 {
				return distance / 2
			} else {
				return distance/2 + 1
			}
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

func (m *Map) SetAt(coordinates Coordinates, value rune) {
	(*m)[coordinates.Y][coordinates.X] = value
}

func (m *Map) IsOutOfBounds(c Coordinates) bool {
	// If row is less than 0
	if c.Y < 0 {
		return true
	}

	// If column is less than 0
	if c.X < 0 {
		return true
	}

	// If row is greater than image length
	if c.Y > len(*m)-1 {
		return true
	}

	// If column is greater than image length
	if c.X > len((*m)[0])-1 {
		return true
	}

	return false
}

func (m *Map) ToLines() []string {
	lines := []string{}

	for _, row := range *m {
		lines = append(lines, string(row))
	}

	return lines
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

	// fmt.Printf("%v%v%v", colorRed, v, colorNone)
}

func printYellow(v any) {
	const colorYellow = "\033[0;33m"
	const colorNone = "\033[0m"

	// fmt.Printf("%v%v%v", colorYellow, v, colorNone)
}

func (m *Map) Amato() []Coordinates {
	capture := false
	hasCaptured := false
	xy := []Coordinates{}

	for rowIndex, row := range *m {
		capture = false
		hasCaptured = false

		for columnIndex, tile := range row {
			// If the tile is a wall.
			if tileIsWall(tile) {
				if capture {
					capture = false
					continue
				}

				if capture == false &&
					(columnIndex > 0 &&
						!tileIsWall(m.At(Coordinates{X: columnIndex - 1, Y: rowIndex}))) {
					capture = true

					// fmt.Printf("\nx:%v y:%v tile:%v capture:%v hasCaptured:%v", columnIndex, rowIndex, string(tile), capture, hasCaptured)

					continue
				}

				// If we've captured already, and
				// we're still capturing when hitting
				// this wall, then stop the capture.
				if hasCaptured && capture {
					capture = false

					// fmt.Printf("\nx:%v y:%v tile:%v capture:%v hasCaptured:%v", columnIndex, rowIndex, string(tile), capture, hasCaptured)

					continue
				}
			} else {
				if capture {
					xy = append(xy, Coordinates{X: columnIndex, Y: rowIndex})
					hasCaptured = true
				}
			}

			// fmt.Printf("\nx:%v y:%v tile:%v capture:%v hasCaptured:%v", columnIndex, rowIndex, string(tile), capture, hasCaptured)

		}
	}

	return xy
}

func tileIsWall(tile rune) bool {
	return tile != '.'
}
