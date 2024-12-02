package lib

import (
	"fmt"
	"math"
	"regexp"
	"slices"
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

func (m *Map) FindFarthestTile(startingPoint Coordinates) (int, []WalkedCell) {
	walked := []WalkedCell{}

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

		walked = append(walked, WalkedCell{Coordinates: currentPosition, Value: m.At(currentPosition)})

		if (distance > 1) && m.At(currentPosition) == 'S' {
			if (distance % 2) == 0 {
				return distance / 2, walked
			}

			return distance/2 + 1, walked
		}

		distance++

		if distance > len(*m)*len((*m)[0]) {
			panic("This went too far!")
		}
	}

	return distance, walked
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

func (m *Map) FindEnclosedTiles() int {
	notNeeded := regexp.MustCompile(`F-*7|L-*J`)

	for rowIndex, row := range *m {
		if notNeeded.MatchString(string(row)) {
			(*m)[rowIndex] = []rune(notNeeded.ReplaceAllString(string(row), " "))
		}

	}

	return 0
}

func (m *Map) Iterate(callback func(c Coordinates)) {
	for rowIndex, row := range *m {
		for columnIndex, _ := range row {
			callback(Coordinates{X: columnIndex, Y: rowIndex})
		}
	}
}

// -----------------------------------------------------------------------------

// isWalkable checks if a cell can be walked on.
func (m *Map) isWalkable(c Coordinates) bool {
	value := m.At(c)

	return slices.Contains(WallTiles, value) && !m.IsOutOfBounds(c)
}

type WalkedCell struct {
	Coordinates Coordinates
	Value       rune
}

func (m *Map) WalkBorders(start Coordinates) []WalkedCell {
	var walkedpath []WalkedCell

	step := 0

	// walk the maze.
	var walk func(coords Coordinates)
	walk = func(coords Coordinates) {
		println("step", step, " - ", coords.X, coords.Y, string(m.At(coords)))
		step++

		if !m.isWalkable(coords) {
			println("not walkable", coords.X, coords.Y, string(m.At(coords)))

			return
		}

		// Add current cell to the path
		if !slices.ContainsFunc(walkedpath, func(wc WalkedCell) bool {
			return wc.Coordinates.SameAs(coords)
		}) {
			walkedpath = append(walkedpath, WalkedCell{Coordinates: coords, Value: m.At(coords)})
		}

		// Explore neighbors
		for _, d := range directions {
			walk(Coordinates{X: coords.X + d.dx, Y: coords.Y + d.dy})
		}
	}

	// Start walking from the starting position
	walk(start)

	return walkedpath
}

// Direction vectors for moving in the maze.
var directions = []struct {
	dx, dy int
}{
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
	{-1, 0}, // Up
}

// CountBoundaryAndInterior calculates the boundary and interior points.
func (m *Map) CalculateArea(path []WalkedCell) int {
	pathvertices := []WalkedCell{}
	for _, cell := range path {
		if slices.Contains(CornerTiles, cell.Value) {
			pathvertices = append(pathvertices, cell)
		}
	}

	return int(CalcPolygonArea(pathvertices))
}

// ShoelaceFormula calculates the total area of a polygon given its vertices.
func ShoelaceFormula(vertices []WalkedCell) float64 {
	n := len(vertices)
	if n < 3 {
		return 0 // Not a polygon
	}

	area := 0
	for i := 0; i < n; i++ {
		// Current vertex
		x1, y1 := vertices[i].Coordinates.X, vertices[i].Coordinates.Y

		// Next vertex (wrap around at the end)
		x2, y2 := vertices[(i+1)%n].Coordinates.X, vertices[(i+1)%n].Coordinates.Y

		area += x1*y2 - y1*x2
	}

	return math.Abs(float64(area)) / 2
}

func CalcPolygonArea(vertices []WalkedCell) float64 {
	total := float64(0)
	l := len(vertices)

	for i := 0; i < l; i++ {
		addX := float64(vertices[i].Coordinates.X)

		yDisplacement := 0
		if i != len(vertices)-1 {
			yDisplacement = i + 1
		}

		xDisplacement := 0
		if i != len(vertices)-1 {
			xDisplacement = i + 1
		}

		addY := float64(vertices[yDisplacement].Coordinates.Y)
		subX := float64(vertices[xDisplacement].Coordinates.X)
		subY := float64(vertices[i].Coordinates.Y)

		total += (addX * addY * float64(0.5))
		total -= (subX * subY * float64(0.5))
	}

	return math.Abs(total)
}

func CalcPolygonArea2(vertices []WalkedCell) float64 {
	total := float64(0)
	l := len(vertices)

	for i := 0; i < l; i++ {
		// Current vertex
		x1 := float64(vertices[i].Coordinates.X)
		y1 := float64(vertices[i].Coordinates.Y)

		// Next vertex (wrap around to the first vertex at the end)
		x2 := float64(vertices[(i+1)%l].Coordinates.X)
		y2 := float64(vertices[(i+1)%l].Coordinates.Y)

		// Shoelace formula components
		total += (x1 * y2) - (y1 * x2)
	}

	return math.Abs(total) / 2
}
