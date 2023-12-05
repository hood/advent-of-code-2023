package day_2_1

import (
	"os"
	"strconv"
	"strings"
)

type CubesSet struct {
	red   int
	green int
	blue  int
}

var AvailableCubesSet = CubesSet{
	red:   12,
	green: 13,
	blue:  14,
}

// Day2.1 entry point.
func Entry() {
	println("\n\n***** Day 2.1 ****")

	fileContent, error := os.ReadFile("./days/day_2_1/input.txt")
	if error != nil {
		panic(error)
	}

	lines := strings.Split(string(fileContent), "\n")

	total := 0

mainLoop:
	for _, line := range lines {
		gameID := parseGameID(line)

		gameCubesSets := parseCubesSets(line)

		for _, partialCubesSet := range gameCubesSets {
			if partialCubesSet.red > AvailableCubesSet.red ||
				partialCubesSet.green > AvailableCubesSet.green ||
				partialCubesSet.blue > AvailableCubesSet.blue {
				continue mainLoop
			}
		}

		total += gameID
	}

	println("Result", "->", total)
}

func parseGameID(line string) int {
	parsedGameID, error := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], "Game ")[1])
	if error != nil {
		panic(error)
	}

	return parsedGameID
}

// Counts the number of cubes sets in a line.
func countCubesSets(line string) int {
	return strings.Count(line, ";") + 1
}

// Counts the number of cubes in a string like "9 red".
func parseCubesCount(stringifiedCount string) int {
	parsedCount, error := strconv.Atoi(strings.Split(stringifiedCount, " ")[0])
	if error != nil {
		panic(error)
	}

	return parsedCount
}

// Given a line, returns the cubes set at the given index.
func parseCubesSet(line string, cubesSetIndex int) CubesSet {
	cubesSet := CubesSet{
		red:   0,
		green: 0,
		blue:  0,
	}

	// Stringified cube set at the given index.
	substring := strings.Split(line, ";")[cubesSetIndex]

	stringifiedCounts := strings.Split(substring, ",")

	for _, stringifiedCount := range stringifiedCounts {
		trimmedStringifiedCount := strings.TrimSpace(stringifiedCount)

		if strings.Contains(trimmedStringifiedCount, "red") {
			cubesSet.red += parseCubesCount(trimmedStringifiedCount)
		}

		if strings.Contains(trimmedStringifiedCount, "green") {
			cubesSet.green += parseCubesCount(trimmedStringifiedCount)
		}

		if strings.Contains(trimmedStringifiedCount, "blue") {
			cubesSet.blue += parseCubesCount(trimmedStringifiedCount)
		}
	}

	return cubesSet
}

func parseCubesSets(line string) []CubesSet {
	cubesSetsCount := countCubesSets(line)
	cubesSets := make([]CubesSet, cubesSetsCount)

	for i := 0; i < cubesSetsCount; i++ {
		cubesSets[i] = parseCubesSet(strings.Split(line, ":")[1], i)
	}

	return cubesSets
}
