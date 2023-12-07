package lib

// Always holds 5 cards.
type Hand struct {
	// Holds information about the presence of each suit.
	PresenceOfFaceValues uint16
	// Holds information about the number of occurrencies of each suit.
	CountOfFaceValues uint16
}

func (h *Hand) AddCard(card rune) {
	h.PresenceOfFaceValues = SetBitAt(h.PresenceOfFaceValues, Cards[card])
	h.CountOfFaceValues += AddBitAt(h.CountOfFaceValues, Cards[card])
}
