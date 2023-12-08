package lib

import (
	"adventofcode2023/days/shared"
	"fmt"
	"testing"
)

func TestCompass(t *testing.T) {
	c := NewCompass()
	c.FromLine("LLR")

	shared.AssertEqual(t, "LLR", string(c.Order))

	expected := []rune{'L', 'L', 'R', 'L', 'L', 'R'}

	for step, expectedRune := range expected {
		direction := c.NextDirection()

		shared.AssertEqual(t, string(expectedRune), string(direction), fmt.Sprintf("Step %v", step))
	}
}
