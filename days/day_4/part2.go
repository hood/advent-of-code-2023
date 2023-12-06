package day_4

import (
	"adventofcode2023/days/day_4/lib"
	"os"
	"strings"
)

// Day4.2 entry point.
func day4Part2() {
	println("\n\n***** Day 4.2 ****")

	fileContent, error := os.ReadFile("./days/day_4/input.txt")
	if error != nil {
		panic(error)
	}

	hands := lib.ExtractHands(strings.Split(string(fileContent), "\n"))

	wonHands := lib.CalculateWonHands(hands)

	result := 0

	for _, won := range wonHands {
		result += won
	}

	println("Result", "->", result)
}
