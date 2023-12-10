package lib

type Map [][]Tile

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
