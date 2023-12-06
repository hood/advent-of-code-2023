package lib

import "strconv"

func ExtractPartsNumbers(lines []string) []PartNumber {
	var partsNumbers []PartNumber
	valueAccumulator := ""

	for row, line := range lines {
		currentPartNumber := PartNumber{
			Value:       -1,
			Row:         row,
			StartColumn: -1,
			EndColumn:   -1,
		}

		for column, char := range line {

			if char >= '0' && char <= '9' {
				// Mark start.
				if currentPartNumber.StartColumn == -1 {
					currentPartNumber.StartColumn = column
				}

				valueAccumulator += string(char)

				continue
			}

			// If was not capturing, continue.
			if currentPartNumber.StartColumn == -1 {
				continue
			}

			// If was capturing, but it's not a number, wrap it up.
			integer, error := strconv.Atoi(valueAccumulator)
			if error != nil {
				panic(error)
			}

			currentPartNumber.EndColumn = column - 1
			currentPartNumber.Value = integer

			partsNumbers = append(partsNumbers, currentPartNumber)

			currentPartNumber = PartNumber{
				Value:       -1,
				Row:         row,
				StartColumn: -1,
				EndColumn:   -1,
			}

			valueAccumulator = ""
		}
	}

	return partsNumbers
}
