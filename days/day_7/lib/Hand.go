package lib

import (
	"sort"
)

const (
	HighCard = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
)

type Hand struct {
	// Always holds 5 cards.
	Cards        []rune
	occurrencies []int
}

func NewHand() *Hand {
	return &Hand{
		Cards:        make([]rune, 15, '-'),
		occurrencies: make([]int, 15),
	}
}

func translateCardValue(card rune) rune {
	switch card {
	case 'T':
		return 10

	case 'J':
		return 11

	case 'Q':
		return 12

	case 'K':
		return 13

	case 'A':
		return 14

	default:
		return card - '0'
	}
}

func translateCardRune(card rune) rune {
	switch card {
	case 'T':
		return 'E'

	case 'J':
		return 'D'

	case 'Q':
		return 'C'

	case 'K':
		return 'B'

	case 'A':
		return 'A'

	default:
		return card
	}
}

func (h *Hand) AddCard(card rune) {
	h.Cards = append(h.Cards, translateCardRune(card))
	h.occurrencies[translateCardValue(card)]++
}

func (h *Hand) NumberScore() int {
	sort.Sort(sort.Reverse(sort.IntSlice(h.occurrencies)))

	switch h.occurrencies[0] {
	case 4:
		return FourOfAKind

	case 3:
		if h.occurrencies[1] == 2 {
			return FullHouse
		}

		return ThreeOfAKind

	case 2:
		if h.occurrencies[1] == 2 {
			return TwoPairs
		}

		return OnePair

	default:
		return HighCard
	}
}

func (h *Hand) Stringified() string {
	result := ""

	for _, card := range h.Cards {
		result += string(card)
	}

	return result
}
