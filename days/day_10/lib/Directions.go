package lib

type Direction struct {
	Y int
	X int
}

var Directions = map[rune]Direction{
	'N': {Y: -1, X: 0},
	'E': {Y: 0, X: 1},
	'S': {Y: 1, X: 0},
	'W': {Y: 0, X: -1},
}

type Tile [4]Direction

var Tiles = map[rune]Tile{
	'|': {Directions['N'], Directions['S']},
	'-': {Directions['W'], Directions['E']},
	'L': {Directions['N'], Directions['E']},
	'J': {Directions['N'], Directions['W']},
	'7': {Directions['S'], Directions['W']},
	'F': {Directions['S'], Directions['E']},
	'S': {Directions['N'], Directions['S'], Directions['W'], Directions['E']},
	'.': {},
}
