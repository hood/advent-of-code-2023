package lib

// Always holds 5 cards.
type Hand struct {
	First, Second, Third, Fourth, Fifth Rank
}

func (h *Hand) isSuited() bool {
	return h.First&h.Second&h.Third&h.Fourth&h.Fifth&0xf000 != 0
}

func (h *Hand) bitMask() int {
	return int(h.First|h.Second|h.Third|h.Fourth|h.Fifth) >> 16
}

func (h *Hand) Score() int {
	if h.isSuited() {
		return Flushes[h.bitMask()]
	}

	if s := Unique5[h.bitMask()]; s != 0 {
		return s
	}

	var (
		k = int((h.First & 0xff) * (h.Second & 0xff) * (h.Third & 0xff) * (h.Fourth & 0xff) * (h.Fifth & 0xff))
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
	case h.First == 0:
		h.First = Cards[card]

	case h.Second == 0:
		h.Second = Cards[card]

	case h.Third == 0:
		h.Third = Cards[card]

	case h.Fourth == 0:
		h.Fourth = Cards[card]

	case h.Fifth == 0:
		h.Fifth = Cards[card]

	default:
		panic("Wtf")
	}

	println(h.Score())
}
