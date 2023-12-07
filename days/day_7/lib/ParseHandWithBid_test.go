package lib

import (
	"adventofcode2023/days/shared"
	"testing"
)

func TestParseHandWithBid(t *testing.T) {
	str := "23456 1"

	handWithBid := ParseHandWithBid(str)

	shared.AssertEqual(t, "23456", "23456")
	shared.AssertEqual(t, 1, handWithBid.Bid)
}
