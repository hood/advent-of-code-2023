package lib

import (
	"adventofcode2023/days/shared"
	"testing"
)

func TestFindInMapWithCompass(t *testing.T) {
	lines := []string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	c := NewCompass()

	c.FromLine(lines[0])
	lines = lines[2:]

	myMap := MapFromLines(lines)

	found, level := FindInMapWithCompass(myMap, c, "ZZZ")

	shared.AssertEqual(t, true, found)
	shared.AssertEqual(t, 6, level)
}
