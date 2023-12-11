package lib

import (
	"fmt"
)

type Map [][]rune

type Coordinates []int

func (c *Coordinates) X() int {
	return (*c)[1]
}

func (c *Coordinates) Y() int {
	return (*c)[0]
}

func MapFromLines(lines []string) (Map, []int) {
	m := Map{}
	startingPoint := []int{0, 0}

	for rowindex, line := range lines {
		row := []rune{}

		for columnindex, char := range line {
			if char == 'S' {
				startingPoint = []int{rowindex, columnindex}
			}

			row = append(row, char)
		}

		m = append(m, row)
	}

	return m, startingPoint
}

func (m *Map) FindFarthestTile(startingPoint Coordinates) int {
	currentX, currentY := startingPoint[0], startingPoint[1]
	previousX, previousY := currentX, currentY

	distance := 0

	// Loop around the current cell.
	// Find all pipes that connect, excluding the previous one.
	for {
		println("\n")
		m.Print([]int{currentX, currentY})

		// Find the connecting tiles.
		connecting := m.FindConnectingTiles(
			Coordinates{currentX, currentY},
			Coordinates{previousX, previousY},
		)

		fmt.Printf("Connecting tiles: %v\n", len(connecting))

		// If there are no connecting tiles, we're done.
		if len(connecting) == 0 {
			break
		}

		for _, connection := range connecting {
			previousX, previousY = currentX, currentY
			currentX, currentY = connection[0], connection[1]
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

	// Check all 4 directions around the point.
	for _, direction := range Directions {
		// West boundary.
		if point.Y() == 0 && direction.Y == -1 {
			continue
		}

		// South boundary.
		if point.Y() == len(*m)-1 && direction.Y == 1 {
			continue
		}

		// North boundary.
		if point.X() == 0 && direction.X == -1 {
			continue
		}

		// East boundary.
		if point.X() == len((*m)[0])-1 && direction.X == 1 {
			continue
		}

		// Calculate the position of the tile in the current direction.
		position := Coordinates{
			point.Y() + direction.Y,
			point.X() + direction.X,
		}

		// Ignore the previous point.
		if position.X() == previous.X() && position.Y() == previous.Y() {
			continue
		}

		// Get the symbol for the calculated position.
		current := m.At(position.X(), position.Y())

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

				break
			}
		}
	}

	return connections
}

func (m *Map) At(x int, y int) rune {
	return (*m)[y][x]
}

func (m *Map) Print(highlight Coordinates) {
	// Clear output.
	// fmt.Print("\033[H\033[2J")

	for rowIndex, row := range *m {
		for columnIndex, tile := range row {
			// for key := range Tiles {
			// 	if tile == key {
			// 		fmt.Printf("%v", key)
			// 	}
			// }

			output := string(tile)

			if rowIndex == highlight[0] && columnIndex == highlight[1] {
				printColored(output)
			} else {
				fmt.Print(output)
			}
		}

		println("")
	}
}

func printColored(v any) {
	const colorRed = "\033[0;31m"
	const colorNone = "\033[0m"

	fmt.Printf("%v%v%v", colorRed, v, colorNone)
}
