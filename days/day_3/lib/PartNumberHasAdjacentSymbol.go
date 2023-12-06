package lib

func PartNumberHasAdjacentSymbol(partNumber PartNumber, symbols []Symbol) bool {
	for _, symbol := range symbols {
		if ValueIsBetween(symbol.Row, partNumber.Row-1, partNumber.Row+1) &&
			(ValueIsBetween(symbol.Column, partNumber.StartColumn-1, partNumber.StartColumn+1) ||
				ValueIsBetween(symbol.Column, partNumber.EndColumn-1, partNumber.EndColumn+1)) {
			return true
		}

	}

	return false
}
