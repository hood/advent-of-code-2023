package day_8

import (
	"adventofcode2023/days/day_8/lib"
	"adventofcode2023/days/shared"
	"strings"
)

func day8Part2() {
	println("\n\n***** Day 8.2 ****")

	lines := shared.ReadFile("./days/day_8/input.txt")

	shared.RunSolution(8, 2, func(callback func(r interface{})) {
		compass := lib.NewCompass()
		compass.FromLine(lines[0])

		// Bump by two because there's a newline separator.
		lines = lines[2:]

		hashMap := lib.HashMapFromLines(lines)

		startingPoints := lib.FindStartingPoints(lines)

		results := make([]int, 0)

		for _, startingPoint := range startingPoints {
			resultFound, resultLevel := lib.FindInHashMapWithCompass(
				hashMap,
				compass,
				startingPoint,
				func(s string) bool { return strings.Contains(s, "Z") },
			)
			if resultFound {
				results = append(results, resultLevel)
			}
		}

		result := lib.LCM(results[0], results[1], results[2:]...)

		callback(result)
	})
}
