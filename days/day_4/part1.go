package day_4

import (
	"adventofcode2023/days/day_4/lib"
	"os"
	"strings"
)

// Day4.1 entry point.
func day4Part1() {
	println("\n\n***** Day 4.1 ****")

	fileContent, error := os.ReadFile("./days/day_4/input.txt")
	if error != nil {
		panic(error)
	}

	hands := lib.ExtractHands(strings.Split(string(fileContent), "\n"))

	result := 0

	for _, hand := range hands {
		result += hand.Points()
	}

	println("Result", "->", result)
}
