package lib

// Always holds 5 cards.
type Hand struct {
	// // Holds information about the presence of each suit.
	// PresenceOfFaceValues uint16
	// // Holds information about the number of occurrencies of each suit.
	// CountOfFaceValues uint16
	First, Second, Third, Fourth, Fifth uint16
}

// func (h *Hand) AddCard(card rune) {
// 	h.PresenceOfFaceValues = SetBitAt(h.PresenceOfFaceValues, Cards[card])
// 	h.CountOfFaceValues += AddBitAt(h.CountOfFaceValues, Cards[card])
// }

// func (h *Hand) Score() uint16 {
// 	// Check for a straight by transforming the first bit field and seeing if it equals binary 11111.
// 	if h.PresenceOfFaceValues&-h.PresenceOfFaceValues == 31 {
// 		return 15
// 	}

// 	return h.CountOfFaceValues modulo 15
// }

func (h *Hand) isSuited() bool {
	return h.First&h.Seecond&h.Third&h.Fourth&h.Fifth&0xf000 != 0
}

func (h *Hand) bitMask() uint16 {
	return h.CountOfFaceValues & 0x1fff
}

func (h *Hand) Score() int {
	if h.isSuited() {
		return Flushes[h.bitMask()]
	}

	// Straights and High Cards
	if s := Unique5[h.Bit()]; s != 0 {
		return s
	}

	// and others... [inlined `findit()`]
	var (
		k = int((h.A & 0xff) * (h.B & 0xff) * (h.C & 0xff) * (h.D & 0xff) * (h.E & 0xff))
	)
	for low, mid, high := 0, 4887>>1, 4887; ; mid = (high + low) >> 1 {
		if product := products[mid]; k < product {
			high = mid - 1
		} else if k > product {
			low = mid + 1
		} else {
			return values[mid]
		}
	}
}
