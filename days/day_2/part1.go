package day_2

import (
	"adventofcode2023/days/day_2/lib"
	"os"
	"strings"
)

var AvailableCubesSet = lib.CubesSet{
	Red:   12,
	Green: 13,
	Blue:  14,
}

// Day2.1 entry point.
func day2Part1() {
	println("\n\n***** Day 2.1 ****")

	fileContent, error := os.ReadFile("./days/day_2/input.txt")
	if error != nil {
		panic(error)
	}

	lines := strings.Split(string(fileContent), "\n")

	total := 0

	for _, line := range lines {
		gameID := lib.ParseGameID(line)

		gameCubesSets := lib.ParseCubesSets(line)

		isFeasible := lib.IsGameFeasible(gameCubesSets, AvailableCubesSet)

		if !isFeasible {
			continue
		}

		total += gameID
	}

	println("Result", "->", total)
}
