package lib

func GearGearAdjacentPartNumbers(row int, column int, partsNumbers []PartNumber) int {
	count := 0
	result := 1

	for _, partNumber := range partsNumbers {
		if ValueIsBetween(row, partNumber.Row-1, partNumber.Row+1) &&
			ValueIsBetween(column, partNumber.StartColumn-1, partNumber.EndColumn+1) {
			count++
			result *= partNumber.Value
		}
	}

	if count < 2 {
		return -1
	}

	return result
}
