package lib

import (
	"adventofcode2023/days/shared"
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

	shared.AssertEqual(t, expectedStringified, resultStringified)
}

func TestSortHandsWithBidsWithEqualScoreHands(t *testing.T) {
	input := []HandWithBid{
		ParseHandWithBid("11111 0"),
		ParseHandWithBid("AAAAA 0"),
		ParseHandWithBid("22222 0"),
		ParseHandWithBid("JJJJJ 0"),
	}

	expected := []HandWithBid{
		ParseHandWithBid("11111 0"),
		ParseHandWithBid("22222 0"),
		ParseHandWithBid("JJJJJ 0"),
		ParseHandWithBid("AAAAA 0"),
	}

	SortHandsWithBids(input)

	expectedStringified := stringifyHands(expected)
	resultStringified := stringifyHands(input)

	shared.AssertEqual(t, expectedStringified, resultStringified)
}

func stringifyHands(hands []HandWithBid) string {
	var builder strings.Builder

	for _, hand := range hands {
		builder.WriteString("\n")
		builder.WriteString(hand.Hand.DebugStringified())
		builder.WriteString("	")
		builder.WriteString(fmt.Sprintf("%v", hand.Hand.NumberScore()))
	}

	return builder.String()
}
