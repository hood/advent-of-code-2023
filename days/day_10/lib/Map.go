package lib

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
	}

	return m, startingPoint
}

// func (m *Map) FindFarthestTile(startingPoint []int) int {
// 	current := Coordinates{startingPoint[0], startingPoint[1]}
// 	previous := Coordinates{startingPoint[0], startingPoint[1]}
// 	distance := 0

// 	// Loop around the current cell.
// 	// Find all pipes that connect, excluding the previous one.
// 	for {
// 		switch {

// 		}
// 	}
// }

func (m *Map) FindConnectingTiles(point Coordinates, previous Coordinates) []Coordinates {
	connections := []Coordinates{}

	for _, direction := range Directions {
		position := Coordinates{point[0] + direction.Y, point[1] + direction.X}

		if position[0] == previous[0] && position[1] == previous[1] {
			continue
		}

		current := (*m)[position[0]][position[1]]

		// Give the endpoints of the current tile, check if they connect to the
		// original tile.
		for _, endpoint := range current {
			if endpoint.Y == direction.Y*-1 && endpoint.X == direction.X*-1 {
				connections = append(connections, position)
			}
		}
	}

	return connections
}
