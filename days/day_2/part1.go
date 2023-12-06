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

mainLoop:
	for _, line := range lines {
		gameID := lib.ParseGameID(line)

		gameCubesSets := lib.ParseCubesSets(line)

		for _, partialCubesSet := range gameCubesSets {
			if partialCubesSet.Red > AvailableCubesSet.Red ||
				partialCubesSet.Green > AvailableCubesSet.Green ||
				partialCubesSet.Blue > AvailableCubesSet.Blue {
				continue mainLoop
			}
		}

		total += gameID
	}

	println("Result", "->", total)
}
