package day_8

import (
	"adventofcode2023/days/day_8/lib"
	"adventofcode2023/days/shared"
	"strings"
)

func day8Part1() {
	lines := shared.ReadFile("./days/day_8/input.txt")

	shared.RunSolution(8, 1, func(callback func(r interface{})) {
		compass := lib.NewCompass()
		compass.FromLine(lines[0])

		// Bump by two because there's a newline separator.
		lines = lines[2:]

		hashMap := lib.HashMapFromLines(lines)

		resultFound, resultLevel := lib.FindInHashMapWithCompass(
			hashMap,
			compass,
			"AAA",
			func(s string) bool { return strings.Contains(s, "ZZZ") },
		)
		if resultFound == false {
			panic("Node `ZZZ` not found!")
		}

		callback(resultLevel)
	})
}
