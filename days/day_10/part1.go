package day_10

import (
	"adventofcode2023/days/day_10/lib"
	"adventofcode2023/days/shared"
)

func Day10Part1() {
	println("\n\n***** Day 10.1 ****")

	lines := shared.ReadFile("./days/day_10/input.txt")

	shared.RunSolution(
		10,
		1,
		func(callback func(r interface{}),
		) {

			m, start := lib.MapFromLines(lines)

			callback(m.FindFarthestTile(start))
		})
}
