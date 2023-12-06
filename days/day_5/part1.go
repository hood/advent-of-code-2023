package day_5

import (
	"adventofcode2023/days/day_5/lib"
	"os"
	"strings"
)

func day5Part1() {
	println("\n\n***** Day 5.1 ****")

	fileContent, error := os.ReadFile("./days/day_5/input.txt")
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

	seeds := lib.ParseSeeds(lines)

	humiditiesByLocations := [][]int{}
	for _, rangeMap := range maps["humidity-to-location"] {
		humiditiesByLocations = append(humiditiesByLocations, lib.GetPairs(rangeMap)...)
	}

	temperaturesByHumidities := [][]int{}
	for _, rangeMap := range maps["temperature-to-humidity"] {
		temperaturesByHumidities = append(temperaturesByHumidities, lib.GetPairs(rangeMap)...)
	}

	lightsByTemperatures := [][]int{}
	for _, rangeMap := range maps["light-to-temperature"] {
		lightsByTemperatures = append(lightsByTemperatures, lib.GetPairs(rangeMap)...)
	}

	watersByLights := [][]int{}
	for _, rangeMap := range maps["water-to-light"] {
		watersByLights = append(watersByLights, lib.GetPairs(rangeMap)...)
	}

	fertilizersByWaters := [][]int{}
	for _, rangeMap := range maps["fertilizer-to-water"] {
		fertilizersByWaters = append(fertilizersByWaters, lib.GetPairs(rangeMap)...)
	}

	soilsByFertilizers := [][]int{}
	for _, rangeMap := range maps["soil-to-fertilizer"] {
		soilsByFertilizers = append(soilsByFertilizers, lib.GetPairs(rangeMap)...)
	}

	seedsBySoils := [][]int{}
	for _, rangeMap := range maps["seed-to-soil"] {
		seedsBySoils = append(seedsBySoils, lib.GetPairs(rangeMap)...)
	}

	lowestSeedLocation := -1

	for index, seed := range seeds {
		seedLocation := lib.GetLocationBySeed(
			seed,
			humiditiesByLocations,
			temperaturesByHumidities,
			lightsByTemperatures,
			watersByLights,
			fertilizersByWaters,
			soilsByFertilizers,
			seedsBySoils,
		)

		if lowestSeedLocation == -1 || seedLocation < lowestSeedLocation {
			lowestSeedLocation = seedLocation
		}

		println("Seed", index, "of", len(seeds), "mapped")
	}

	println("Result", "->", lowestSeedLocation)
}
