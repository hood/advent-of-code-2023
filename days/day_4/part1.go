package day_4

import (
	"adventofcode2023/days/day_4/lib"
	"fmt"
	"os"
	"strings"
)

// Day4.1 entry point.
func day4Part1() {
	println("\n\n***** Day 4.1 ****")

	fileContent, error := os.ReadFile("./days/day_4/part_1_test_input.txt")
	if error != nil {
		panic(error)
	}

	hands := lib.ExtractHands(strings.Split(string(fileContent), "\n"))

	for _, hand := range hands {
		println("Hand", hand.ID, "->", fmt.Sprintf("%v", hand.WinningNumbers()))
	}

	// println("Result", "->", fileContent)
}
