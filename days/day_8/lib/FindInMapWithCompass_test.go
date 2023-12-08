package lib

import (
	"adventofcode2023/days/shared"
	"strings"
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

	found, level := FindInHashMapWithCompass(myMap, c, "ZZZ")

	shared.AssertEqual(t, true, found)
	shared.AssertEqual(t, 6, level)
}

func TestFindInHashMapWithDoublePath(t *testing.T) {
	lines := []string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}

	c := NewCompass()

	c.FromLine(lines[0])
	lines = lines[2:]

	myMap := HashMapFromLines(lines)

	startingPoints := FindStartingPoints(lines)

	found, level := FindInHashMapWithDoublePath(
		myMap,
		c,
		startingPoints,
		func(v string) bool {
			return strings.Contains(v, "Z")
		},
	)

	shared.AssertEqual(t, true, found)
	shared.AssertEqual(t, 6, level)
}
