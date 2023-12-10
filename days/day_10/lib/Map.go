package lib

import "fmt"

type Map [][]Tile

type Coordinates []int

func MapFromLines(lines []string) (Map, []int) {
	m := Map{}
	startingPoint := []int{0, 0}

	for rowindex, line := range lines {
		row := []Tile{}

		for columnindex, char := range line {
			if char == 'S' {
				startingPoint = []int{rowindex, columnindex}
			}

			row = append(row, Tiles[char])
		}

		m = append(m, row)
	}

	return m, startingPoint
}

func (m *Map) FindFarthestTile(startingPoint []int) int {
	currentX, currentY := startingPoint[0], startingPoint[1]
	previousX, previousY := startingPoint[0], startingPoint[1]
	distance := 0

	// Loop around the current cell.
	// Find all pipes that connect, excluding the previous one.
	for {
		connecting := m.FindConnectingTiles([]int{currentX, currentY}, []int{previousX, previousY})

		fmt.Printf("x, y: %v, %v\nConnecting tiles: %v\n", currentX, currentY, connecting)

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
func (m *Map) FindConnectingTiles(point Coordinates, previous Coordinates) []Coordinates {
	connections := []Coordinates{}

	// Check all directions around the point.
	for _, direction := range Directions {
		// West boundary.
		if point[0] == 0 && direction.X == -1 {
			continue
		}

		// East boundary.
		if point[0] == len((*m)[0])-1 && direction.X == 1 {
			continue
		}

		// North boundary.
		if point[1] == 0 && direction.Y == -1 {
			continue
		}

		// South boundary.
		if point[1] == len(*m)-1 && direction.Y == 1 {
			continue
		}

		position := Coordinates{point[0] + direction.X, point[1] + direction.Y}

		// Ignore the previous point.
		if position[0] == previous[0] && position[1] == previous[1] {
			continue
		}

		current := m.At(point[0], point[1])

		// Loop the endpoints of the current tile, check if they connect to the
		// original tile.
		for _, endpoint := range current {
			if endpoint.X == -direction.X &&
				endpoint.Y == -direction.Y {
				connections = append(connections, position)

				break
			}
		}
	}

	return connections
}

func (m *Map) At(x int, y int) Tile {
	return (*m)[y][x]
}

func (m *Map) Print(highlight Coordinates) {
	for rowIndex, row := range *m {
		for columnIndex, tile := range row {
			for key, endpoints := range Tiles {
				if tile == endpoints {
					fmt.Printf("%v", key)
				}
			}

			output := ""

			if rowIndex == highlight[0] && columnIndex == highlight[1] {
				output = fmt.Sprintf("\x1b[%dm%v\x1b[0m", 34, tile)
			} else {
				output = fmt.Sprintf("%v", tile)
			}

			fmt.Print(output)
		}

		fmt.Println()
	}
}
