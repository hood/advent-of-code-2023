package lib

// Always holds 5 cards.
type Hand struct {
	Cards []CardInfo
}

func NewHand() *Hand {
	return &Hand{
		Cards: make([]CardInfo, 0),
	}
}

func (h *Hand) AddCard(card rune) {
	h.Cards = append(h.Cards, CardsInfo[card])
}

func (h *Hand) NumberScore() int {
	result := 1

	for _, card := range h.Cards {
		result *= card.Value
	}

	return result
}

func (h *Hand) StringScore(numberScore uint64) string {
	switch {
	case numberScore > 6185:
		return "high-card"

	case numberScore > 3325:
		return "one-pair"

	case numberScore > 2467:
		return "two-pairs"

	case numberScore > 1609:
		return "three-of-a-kind"

	case numberScore > 1599:
		return "straight"

	case numberScore > 322:
		return "flush"

	case numberScore > 166:
		return "full-house"

	case numberScore > 10:
		return "four-of-a-kind"

	default:
		return "straight-flush"
	}
}

func (h *Hand) Stringified() string {
	result := ""

	for _, card := range h.Cards {
		result += string(card.Symbol)
	}

	return result
}
