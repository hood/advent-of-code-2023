package lib

func ExtractGears(lines []string, partsNumbers []PartNumber) []Gear {
	gears := []Gear{}

	for row, line := range lines {
		for column, char := range line {
			if char != '*' {
				continue
			}

			ratio := GearGearAdjacentPartNumbers(row, column, partsNumbers)

			if ratio == -1 {
				continue
			}

			gears = append(gears, Gear{
				Row:    row,
				Column: column,
				Ratio:  ratio,
			})
		}
	}

	return gears
}
