package day_3

import (
	"adventofcode2023/days/day_3/lib"
	"os"
	"strings"
)

// Day3.2 entry point.
func day3Part2() {
	println("\n\n***** Day 3.2 ****")

	fileContent, error := os.ReadFile("./days/day_3/input.txt")
	if error != nil {
		panic(error)
	}

	lines := strings.Split(string(fileContent), "\n")

	partsNumbers := lib.ExtractPartsNumbers(lines)
	gears := lib.ExtractGears(lines, partsNumbers)

	result := 0

	for _, gear := range gears {
		result += gear.Ratio
	}

	println("Result", "->", result)
}
