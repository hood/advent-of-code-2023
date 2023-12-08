package lib

import (
	"adventofcode2023/days/shared"
	"testing"
)

func TestCompass(t *testing.T) {
	c := NewCompass()
	c.FromLine("LLR")

	expected := []rune{'L', 'L', 'R', 'L', 'L', 'R'}

	for _, expectedRune := range expected {
		next := c.NextDirection()

		shared.AssertEqual(t, string(expectedRune), string(next))
	}
}
