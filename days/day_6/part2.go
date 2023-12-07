package day_6

import (
	"adventofcode2023/days/day_6/lib"
	"adventofcode2023/days/shared"
)

func day6Part2() {
	println("\n\n***** Day 6.2 ****")

	lines := shared.ReadFile("./days/day_6/input.txt")

	race := lib.ParseRace(lines)

	result := len(lib.FindWaysToWinRace(&race))

	println("Result", "->", result)
}
