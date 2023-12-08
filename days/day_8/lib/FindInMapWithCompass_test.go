package lib

import (
	"adventofcode2023/days/shared"
	"testing"
)

func TestFindInHashMapWithCompass(t *testing.T) {
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

	myMap := HashMapFromLines(lines)

	found, level := FindInHashMapWithCompass(myMap, c, "AAA", func(s string) bool { return s == "ZZZ" })

	shared.AssertEqual(t, true, found)
	shared.AssertEqual(t, 6, level)
}
