package lib

import (
	"fmt"
	"strings"
	"testing"
)

func TestSortHandsWithBids(t *testing.T) {
	input := []HandWithBid{
		ParseHandWithBid("KKKK9 0"),
		ParseHandWithBid("KKK99 0"),
		ParseHandWithBid("KK992 0"),
		ParseHandWithBid("AJ942 0"),
	}

	expected := []HandWithBid{
		ParseHandWithBid("AJ942 0"),
		ParseHandWithBid("KK992 0"),
		ParseHandWithBid("KKK99 0"),
		ParseHandWithBid("KKKK9 0"),
	}

	SortHandsWithBids(input)

	expectedStringified := stringifyHands(expected)
	resultStringified := stringifyHands(input)

	if resultStringified != expectedStringified {
		t.Errorf("\nGot%v\n\nExpected%v", resultStringified, expectedStringified)
	}
}

func stringifyHands(hands []HandWithBid) string {
	var builder strings.Builder

	for _, hand := range hands {
		builder.WriteString("\n")
		builder.WriteString(hand.Hand.Stringified())
		builder.WriteString("	")
		builder.WriteString(fmt.Sprintf("%v", hand.Hand.NumberScore()))
	}

	return builder.String()
}
