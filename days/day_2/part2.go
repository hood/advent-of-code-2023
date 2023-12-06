package day_2

import (
	"adventofcode2023/days/day_2/lib"
	"os"
	"strings"
)

func day2Part2() {
	println("\n\n***** Day 2.2 ****")

	fileContent, error := os.ReadFile("./days/day_2/input.txt")
	if error != nil {
		panic(error)
	}

	games := strings.Split(string(fileContent), "\n")

	total := 0

	for _, game := range games {
		gameCubesSets := lib.ParseCubesSets(game)

		result := lib.CountMinimumNeededGameCubesSet(gameCubesSets)

		power := result.Red * result.Green * result.Blue

		total += power
	}

	println("Result", "->", total)
}
