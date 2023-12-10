package day_9

import (
	"adventofcode2023/days/day_9/lib"
	"adventofcode2023/days/shared"
)

func Day9Part1() {
	lines := shared.ReadFile("./days/day_9/input.txt")

	shared.RunSolution(9, 1, func(callback func(r interface{})) {
		result := 0

		for _, line := range lines {
			s := lib.SequenceFromString(line)
			// result += lib.Lagrange(s.Values)
			result += s.FinalValue()
		}

		callback(result)
	})
}
