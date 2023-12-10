package day_9

import (
	"adventofcode2023/days/day_9/lib"
	"adventofcode2023/days/shared"
	"testing"
)

func TestDay9Part1Test(t *testing.T) {
	lines := shared.ReadFile("./test_input.txt")

	result := 0

	for _, line := range lines {
		s := lib.SequenceFromString(line)

		result += lib.Lagrange(s.Values)
	}

	shared.AssertEqual(t, 114, result)
}

func TestDay9Part1Final(t *testing.T) {
	lines := shared.ReadFile("./input.txt")

	result := 0

	for _, line := range lines {
		s := lib.SequenceFromString(line)

		result += s.FinalValue()
	}

	shared.AssertEqual(t, 1861775706, result)
}
