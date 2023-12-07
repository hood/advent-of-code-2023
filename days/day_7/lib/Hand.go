package lib

import "fmt"

// Always holds 5 cards.
type Hand struct {
	// Cards []Rank
	// Bit mask that represents the ranks of the cards in the hand.
	ranksPresences uint16
	// Bit mask that represents the number of times a rank appears in the hand.
	ranksOccurrencies uint16
}

func NewHand() *Hand {
	return &Hand{
		ranksPresences:    0b0000000000000000,
		ranksOccurrencies: 0b0000000000000000,
	}
}

func (h *Hand) AddCard(card rune) {
	// switch {
	// case len(h.Cards) < 5:
	// h.Cards = append(h.Cards, Cards[card])
	position := uint8(Cards[card])

	println("position", position)

	h.ranksPresences = SetBitAt(h.ranksPresences, position)
	h.ranksOccurrencies = AddBitAt(h.ranksOccurrencies, position)
	// h.ranksPresences |= 1 << byte(Cards[card])
	// h.ranksOccurrencies += 1 << byte(Cards[card])

	// default:
	// panic("Wtf")
	// }
}

func (h *Hand) Score() uint16 {
	fmt.Printf("ranksPresences\n"+"%016b\n", h.ranksPresences)
	fmt.Printf("ranksOccurrencies\n"+"%016b\n", h.ranksOccurrencies)
	return 1
}

func printBitmask(bitmask uint16) {

}

// 15 times 0
// 0000000000000000
