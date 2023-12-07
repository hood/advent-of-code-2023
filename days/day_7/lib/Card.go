package lib

type Rank uint8

// const (
// 	Deuce Rank = iota
// 	Trey
// 	Four
// 	Five
// 	Six
// 	Seven
// 	Eight
// 	Nine
// 	Ten
// 	Jack
// 	Queen
// 	King
// 	Ace
// )

// const (
// 	Deuce Rank = 12
// 	Trey  Rank = 11
// 	Four  Rank = 10
// 	Five  Rank = 9
// 	Six   Rank = 8
// 	Seven Rank = 7
// 	Eight Rank = 6
// 	Nine  Rank = 5
// 	Ten   Rank = 4
// 	Jack  Rank = 3
// 	Queen Rank = 2
// 	King  Rank = 1
// 	Ace   Rank = iota
// )

const (
	Deuce Rank = 3
	Trey  Rank = 4
	Four  Rank = 5
	Five  Rank = 6
	Six   Rank = 7
	Seven Rank = 8
	Eight Rank = 9
	Nine  Rank = 10
	Ten   Rank = 11
	Jack  Rank = 12
	Queen Rank = 13
	King  Rank = 14
	Ace   Rank = 15
)

var Cards = map[rune]Rank{
	'A': Ace,
	'K': King,
	'Q': Queen,
	'J': Jack,
	'T': Ten,
	'9': Nine,
	'8': Eight,
	'7': Seven,
	'6': Six,
	'5': Five,
	'4': Four,
	'3': Trey,
	'2': Deuce,
}
