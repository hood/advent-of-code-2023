package lib

type Direction struct {
	Vertical   int
	Horizontal int
}

var Directions = map[rune]Direction{
	'N': Direction{Vertical: -1, Horizontal: 0},
	'E': Direction{Vertical: 0, Horizontal: 1},
	'S': Direction{Vertical: 1, Horizontal: 1},
	'W': Direction{Vertical: 0, Horizontal: -1},
}

type Tile []Direction

var Tiles = map[rune]Tile{
	'|': {Directions['N'], Directions['S']},
	'-': {Directions['W'], Directions['E']},
	'L': {Directions['N'], Directions['E']},
	'J': {Directions['N'], Directions['W']},
	'7': {Directions['S'], Directions['W']},
	'F': {Directions['S'], Directions['E']},
	'S': {Directions['N'], Directions['S'], Directions['W'], Directions['E']},
}
