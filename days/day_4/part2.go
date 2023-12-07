package day_4

import (
	"adventofcode2023/days/day_4/lib"
	"adventofcode2023/days/shared"
)

// Day4.2 entry point.
func day4Part2() {
	println("\n\n***** Day 4.2 ****")

	hands := lib.ExtractHands(shared.ReadFile(("./days/day_4/input.txt")))

	wonHands := lib.CalculateWonHands(hands)

	result := 0

	for _, won := range wonHands {
		result += won
	}

	println("Result", "->", result)
}
