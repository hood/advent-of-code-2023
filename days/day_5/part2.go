package day_5

import (
	"adventofcode2023/days/day_5/lib"
	"os"
	"strings"
)

func day5Part2() {
	println("\n\n***** Day 5.2 ****")

	fileContent, error := os.ReadFile("./days/day_5/part_1_test_input.txt")
	if error != nil {
		panic(error)
	}

	lines := strings.Split(string(fileContent), "\n")

	maps := map[string][]lib.Map{
		"seed-to-soil":            {},
		"soil-to-fertilizer":      {},
		"fertilizer-to-water":     {},
		"water-to-light":          {},
		"light-to-temperature":    {},
		"temperature-to-humidity": {},
		"humidity-to-location":    {},
	}

	for mapName := range maps {
		mapsGroup := lib.GetMapsGroup(lines, mapName)

		maps[mapName] = lib.ParseMaps(mapsGroup)
	}

	seedsRanges := lib.ParseSeedsRanges(lines)

	lowestSeedLocation := -1

	for _, seedsRange := range seedsRanges {
		seedLocation := lib.GetLocationBySeedsRange(
			seedsRange,
			maps["seed-to-soil"],
			maps["soil-to-fertilizer"],
			maps["fertilizer-to-water"],
			maps["water-to-light"],
			maps["light-to-temperature"],
			maps["temperature-to-humidity"],
			maps["humidity-to-location"],
		)

		if lowestSeedLocation == -1 || seedLocation < lowestSeedLocation {
			lowestSeedLocation = seedLocation
		}
	}

	println("Result", "->", lowestSeedLocation)
}

// your answer is too high: 289863851
