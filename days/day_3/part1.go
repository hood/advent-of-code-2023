package day_3

import (
	"adventofcode2023/days/day_3/lib"
	"os"
	"strings"
)

// Day3.1 entry point.
func day3Part1() {
	println("\n\n***** Day 3.1 ****")

	fileContent, error := os.ReadFile("./days/day_3/test_input.txt")
	if error != nil {
		panic(error)
	}

	lines := strings.Split(string(fileContent), "\n")

	partsNumbers := lib.ExtractPartsNumbers(lines)
	symbols := lib.ExtractSymbols(lines)

	result := 0

	for _, partNumber := range partsNumbers {
		if lib.PartNumberHasAdjacentSymbol(partNumber, symbols) {
			result += partNumber.Value
		}
	}

	println("Result", "->", result)
}
