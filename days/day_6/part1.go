package day_6

import (
	"adventofcode2023/days/day_6/lib"
	"adventofcode2023/days/shared"
)

func day6Part1() {
	println("\n\n***** Day 6.1 ****")

	lines := shared.ReadFile("./days/day_6/input.txt")

	races := lib.ParseRaces(lines)

	result := 1

	for _, race := range races {
		for _, win := range lib.FindWaysToWinRace(&race) {
			result *= win
		}
	}

	println("Result", "->", result)
}
