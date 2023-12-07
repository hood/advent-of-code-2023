package lib

import (
	"adventofcode2023/days/shared"
	"strings"
)

func ParseHandWithBid(str string) HandWithBid {
	strs := strings.Split(str, "")

	cards := strs[0]
	bid := shared.ParseInteger(strs[1])

	hand := NewHand()

	for _, card := range cards {
		hand.AddCard(card)
	}

	return HandWithBid{
		Hand: *hand,
		Bid:  bid,
	}
}
