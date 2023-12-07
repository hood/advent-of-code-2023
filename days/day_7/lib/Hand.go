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

func (h *Hand) IsStraight() bool {
	// [5, 6, 7, 8, 9] translates to binary 000001111100000.
	// We use bitwise AND with the negative of itself to get the least significant bit: 000001111100000 & -000001111100000 = 100000
	// We divide the original bit field by the LSB which normalizes it: 000001111100000 / 100000 = 11111

	binary := h.ranksOccurrencies
	lsb := binary & -binary
	normalized := binary / lsb

	return normalized == 0b11111
}

func (h *Hand) Score() uint64 {
	h.DebugPresences()
	h.DebugOccurrencies()

	return uint64(h.ranksOccurrencies) % 15
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
