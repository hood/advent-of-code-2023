package day_5

import (
	"adventofcode2023/days/day_5/lib"
	"os"
	"strings"
)

func day5Part1() {
	println("\n\n***** Day 5.1 ****")

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

}
