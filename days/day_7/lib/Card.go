package lib

type Rank uint16

const (
	Deuce Rank = iota
	Trey
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
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
