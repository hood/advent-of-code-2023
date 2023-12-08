package day_7

import (
	"adventofcode2023/days/day_7/lib"
	"adventofcode2023/days/shared"
	"sort"
)

func day7Part1() {
	println("\n\n***** Day 7.1 ****")

	lines := shared.ReadFile("./days/day_7/input.txt")

	shared.RunSolution(func(callback func(r interface{})) {
		handsWithBids := make([]lib.HandWithBid, len(lines))

		for i, line := range lines {
			handsWithBids[i] = lib.ParseHandWithBid(line)
		}

		sort.Sort(lib.HandsSorter(handsWithBids))

		result := 0

		for index, handWithBid := range handsWithBids {
			println(handWithBid.Hand.DebugStringified())

			result += handWithBid.Bid * (index + 1)
		}

		callback(result)
	})
}
