package lib

func CharIsSymbol(char rune) bool {
	return char == '#' ||
		char == '*' ||
		char == '+' ||
		char == '$' ||
		char == '@'
}
