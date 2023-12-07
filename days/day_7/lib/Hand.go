package lib

import "fmt"

// Always holds 5 cards.
type Hand struct {
	// Cards []Rank
	// Bit mask that represents the ranks of the cards in the hand.
	ranksPresences uint16
	// Bit mask that represents the number of times a rank appears in the hand.
	ranksOccurrencies uint64
}

func NewHand() *Hand {
	return &Hand{
		ranksPresences:    0b0000000000000000,
		ranksOccurrencies: 0b0000000000000000000000000000000000000000000000000000000000000,
	}
}

func (h *Hand) AddCard(card rune) {
	position := uint8(Cards[card])

	h.ranksPresences = SetBitAt(h.ranksPresences, position)
	h.ranksOccurrencies = AddBitAt(h.ranksOccurrencies, position)

	h.DebugOccurrencies()
}

func (h *Hand) Score() uint16 {
	h.DebugPresences()
	h.DebugOccurrencies()
	return 1
}

func (h *Hand) DebugPresences() {
	str := fmt.Sprintf("%016b\n", h.ranksPresences)
	result := ""
	infoRow := "A | K | Q | J | T | 9 | 8 | 7 | 6 | 5 | 4 | 3 | 2\n"

	for i := 0; i < len(str); i++ {
		if i%1 == 0 && i != 0 {
			result += " | "
		}

		result += string(str[i])
	}

	println("PRESENCES:\n", result, infoRow)
}

func (h *Hand) DebugOccurrencies() {
	str := fmt.Sprintf("%064b\n", h.ranksOccurrencies)
	result := ""
	infoRow := "A    | K    | Q    | J    | T    | 9    | 8    | 7    | 6    | 5    | 4    | 3    | 2\n"

	for i := 0; i < len(str); i++ {
		if i%4 == 0 && i != 0 {
			result += " | "
		}

		result += string(str[i])
	}

	println("OCCURRENCIES:\n", result, infoRow)
}
