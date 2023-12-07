package lib

type Rank uint8

type CardInfo struct {
	Symbol rune
	Rank   int
	Value  int
}

var CardsInfo = map[rune]CardInfo{
	'A': {
		Symbol: 'A',
		Rank:   12,
		Value:  41,
	},
	'K': {
		Symbol: 'K',
		Rank:   11,
		Value:  37,
	},
	'Q': {
		Symbol: 'Q',
		Rank:   10,
		Value:  31,
	},
	'J': {
		Symbol: 'J',
		Rank:   9,
		Value:  29,
	},
	'T': {
		Symbol: 'T',
		Rank:   8,
		Value:  23,
	},
	'9': {
		Symbol: '9',
		Rank:   7,
		Value:  19,
	},
	'8': {
		Symbol: '8',
		Rank:   6,
		Value:  17,
	},
	'7': {
		Symbol: '7',
		Rank:   5,
		Value:  13,
	},
	'6': {
		Symbol: '6',
		Rank:   4,
		Value:  11,
	},
	'5': {
		Symbol: '5',
		Rank:   3,
		Value:  7,
	},
	'4': {
		Symbol: '4',
		Rank:   2,
		Value:  5,
	},
	'3': {
		Symbol: '3',
		Rank:   1,
		Value:  3,
	},
	'2': {
		Symbol: '2',
		Rank:   0,
		Value:  2,
	},
}
