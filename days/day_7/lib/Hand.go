package lib

import "fmt"

// Always holds 5 cards.
type Hand struct {
	Cards []Rank
}

func (h *Hand) isSuited() bool {
	return h.Cards[0]&h.Cards[1]&h.Cards[2]&h.Cards[3]&h.Cards[4]&0xf000 != 0
}

func (h *Hand) bitMask() int {
	return int(h.Cards[0]|h.Cards[1]|h.Cards[2]|h.Cards[3]|h.Cards[4]) >> 16
}

func (h *Hand) Score() int {
	if h.isSuited() {
		return Flushes[h.bitMask()]
	}

	if s := Unique5[h.bitMask()]; s != 0 {
		return s
	}

	var (
		k = int((h.Cards[0] & 0xff) * (h.Cards[1] & 0xff) * (h.Cards[2] & 0xff) * (h.Cards[3] & 0xff) * (h.Cards[4] & 0xff))
	)
	for low, mid, high := 0, 4887>>1, 4887; ; mid = (high + low) >> 1 {
		if product := Products[mid]; k < product {
			high = mid - 1
		} else if k > product {
			low = mid + 1
		} else {
			return Values[mid]
		}
	}
}

func (h *Hand) AddCard(card rune) {
	switch {
	case len(h.Cards) < 5:
		h.Cards = append(h.Cards, Cards[card])

	default:
		panic("Wtf")
	}

	println(h.String())
}

func (h Hand) String() string {
	if len(h.Cards) < 5 {
		return "Not enough cards"
	}

	return fmt.Sprintf("%04b %04b %04b %04b %04b", h.Cards[0], h.Cards[1], h.Cards[2], h.Cards[3], h.Cards[4])
}
