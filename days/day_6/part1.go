package day_6

import (
	"adventofcode2023/days/day_6/lib"
	"adventofcode2023/days/shared"
	"fmt"
)

func day6Part1() {
	println("\n\n***** Day 6.1 ****")

	lines := shared.ReadFile("./days/day_6/part_1_test_input.txt")

	races := lib.ParseRaces(lines)

	for _, race := range races {
		lib.FindWaysToWinRace(&race)
	}

	result := 0

	println("Result", "->", fmt.Sprintf("%v", result))
}
