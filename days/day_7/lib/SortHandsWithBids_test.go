package lib

import (
	"adventofcode2023/days/shared"
	"fmt"
	"strings"
	"testing"
)

func TestSortHandsWithBids(t *testing.T) {
	input := []HandWithBid{
		ParseHandWithBid("KKKK9 0", false),
		ParseHandWithBid("KKK99 0", false),
		ParseHandWithBid("KK992 0", false),
		ParseHandWithBid("AJ942 0", false),
	}

	expected := []HandWithBid{
		ParseHandWithBid("AJ942 0", false),
		ParseHandWithBid("KK992 0", false),
		ParseHandWithBid("KKK99 0", false),
		ParseHandWithBid("KKKK9 0", false),
	}

	SortHandsWithBids(input)

	expectedStringified := stringifyHands(expected)
	resultStringified := stringifyHands(input)

	shared.AssertEqual(t, expectedStringified, resultStringified)
}

func TestSortHandsWithBidsWithEqualScoreHands(t *testing.T) {
	input := []HandWithBid{
		ParseHandWithBid("11111 0", false),
		ParseHandWithBid("AAAAA 0", false),
		ParseHandWithBid("22222 0", false),
		ParseHandWithBid("JJJJJ 0", false),
	}

	expected := []HandWithBid{
		ParseHandWithBid("11111 0", false),
		ParseHandWithBid("22222 0", false),
		ParseHandWithBid("JJJJJ 0", false),
		ParseHandWithBid("AAAAA 0", false),
	}

	SortHandsWithBids(input)

	expectedStringified := stringifyHands(expected)
	resultStringified := stringifyHands(input)

	shared.AssertEqual(t, expectedStringified, resultStringified)
}

func TestSortHandsWithBidsWithJokers(t *testing.T) {
	input := []HandWithBid{
		ParseHandWithBid("32T3K -1", true),
		ParseHandWithBid("T55J5 -1", true),
		ParseHandWithBid("KK677 -1", true),
		ParseHandWithBid("KTJJT -1", true),
		ParseHandWithBid("QQQJA -1", true),
	}

	expected := []HandWithBid{
		ParseHandWithBid("32T3K -1", true),
		ParseHandWithBid("KK677 -1", true),
		ParseHandWithBid("T55J5 -1", true),
		ParseHandWithBid("QQQJA -1", true),
		ParseHandWithBid("KTJJT -1", true),
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
