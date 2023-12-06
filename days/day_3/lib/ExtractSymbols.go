package lib

func ExtractSymbols(lines []string) []Symbol {
	var symbols []Symbol

	for row, line := range lines {
		for column, char := range line {
			if CharIsSymbol(char) {
				symbols = append(symbols, Symbol{
					Row:    row,
					Column: column,
				})
			}
		}
	}

	return symbols
}
