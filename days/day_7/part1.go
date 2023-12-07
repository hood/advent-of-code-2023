package day_7

import (
	"adventofcode2023/days/day_7/lib"
	"adventofcode2023/days/shared"
)

func day7Part1() {
	println("\n\n***** Day 7.1 ****")

	lines := shared.ReadFile("./days/day_7/part_1_test_input.txt")

	handsWithBids := make([]lib.HandWithBid, len(lines))

	for i, line := range lines {
		handsWithBids[i] = lib.ParseHandWithBid(line)
	}

}
