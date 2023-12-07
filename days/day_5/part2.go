package day_5

import (
	"adventofcode2023/days/day_5/lib"
	"adventofcode2023/days/shared"
	"os"
	"strings"
)

type BySourceRangeStart []lib.Map

func (a BySourceRangeStart) Len() int           { return len(a) }
func (a BySourceRangeStart) Less(i, j int) bool { return a[i].SourceRangeStart < a[j].SourceRangeStart }
func (a BySourceRangeStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func day5Part2() {
	println("\n\n***** Day 5.2 ****")

	fileContent, error := os.ReadFile("./days/day_5/input.txt")
	if error != nil {
		panic(error)
	}

	lines := strings.Split(string(fileContent), "\n")

	maps := map[string]*shared.BinarySearchTree[lib.Map]{
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

	for i, seedsRange := range seedsRanges {

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

		println("Seeds range", i+1, "of", len(seedsRanges), "result", "->", lowestSeedLocation)

	}

	println("Result", "->", lowestSeedLocation)
}
