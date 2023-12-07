package lib

// Always holds 5 cards.
type Hand struct {
	// Holds information about the presence of each suit.
	PresenceOfFaceValues uint16
	// Holds information about the number of occurrencies of each suit.
	CountOfFaceValues uint16
}

func (h *Hand) AddCard(card rune) {
	// h.PresenceOfFaceValues |= 1 << Cards[card]
	// h.CountOfFaceValues += 1 << Cards[card]

	h.PresenceOfFaceValues = SetBitAt(h.PresenceOfFaceValues, Cards[card])
	h.CountOfFaceValues += AddBitAt(h.CountOfFaceValues, Cards[card])
}

// const (
// 	FiveOfAKind  = 7
// 	FourOfAKind  = 6
// 	FullHouse    = 5
// 	ThreeOfAKind = 4
// 	TwoPair      = 3
// 	OnePair      = 2
// 	HighCard     = 1
// )

// type Hand struct {
// 	// 5 cards.
// 	Cards []string
// }

// func (h *Hand) AddCard(card string) {
// 	h.Cards = append(h.Cards, card)
// }

// func (h *Hand) Type() int {
// 	if h.isFiveOfAKind() {
// 		return FiveOfAKind
// 	}

// 	if h.isFourOfAKind() {
// 		return FourOfAKind
// 	}

// 	if h.isFullHouse() {
// 		return FullHouse
// 	}

// 	if h.isThreeOfAKind() {
// 		return ThreeOfAKind
// 	}

// 	if h.isTwoPair() {
// 		return TwoPair
// 	}

// 	if h.isOnePair() {
// 		return OnePair
// 	}

// 	return HighCard

// }

// func (h *Hand) isFiveOfAKind() bool {
// 	return h.isFourOfAKind() && h.isOnePair()
// }

// func (h *Hand) isFourOfAKind() bool {
// 	return h.isThreeOfAKind() && h.isOnePair()
// }

// func (h *Hand) isFullHouse() bool {
// 	return h.isThreeOfAKind() && h.isTwoPair()
// }

// func (h *Hand) isThreeOfAKind() bool {
// 	return h.isOnePair() && h.isOnePair()
// }

// func (h *Hand) isTwoPair() bool {
// 	return h.isOnePair() && h.isOnePair()
// }

// func (h *Hand) isOnePair() bool {
// 	//
// }

// func (h *Hand) isHighCard() bool {

// }

// // Every hand is exactly one type. From strongest to weakest, they are:

// // Five of a kind, where all five cards have the same label: AAAAA
// // Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// // Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// // Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// // Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// // One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// // High card, where all cards' labels are distinct: 23456
