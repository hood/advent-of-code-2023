package day_8

import (
	"adventofcode2023/days/day_8/lib"
	"adventofcode2023/days/shared"
	"strings"
)

func day8Part2() {
	println("\n\n***** Day 8.2 ****")

	lines := shared.ReadFile("./days/day_8/input.txt")

	shared.RunSolution(func(callback func(r interface{})) {
		compass := lib.NewCompass()
		compass.FromLine(lines[0])

		// Bump by two because there's a newline separator.
		lines = lines[2:]

		hashMap := lib.HashMapFromLines(lines)

		startingPoints := lib.FindStartingPoints(lines)

		resultFound, resultLevel := lib.FindInHashMapWithDoublePath(hashMap, compass, startingPoints, func(s string) bool {
			return strings.Contains(s, "Z")
		})
		if resultFound == false {
			panic("End not found!")
		}

		callback(resultLevel)
	})
}
