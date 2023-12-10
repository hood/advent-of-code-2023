package day_9

import (
	"adventofcode2023/days/day_9/lib"
	"adventofcode2023/days/shared"
	"testing"
)

func TestDay9Part2Test(t *testing.T) {
	lines := shared.ReadFile("./test_input.txt")

	result := 0

	for _, line := range lines {
		s := lib.SequenceFromString(line)

		result += s.FinalReversedValue()
	}

	shared.AssertEqual(t, 2, result)
}
