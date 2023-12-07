package lib

import (
	"adventofcode2023/days/shared"
)

func GetLocationBySeedsRange(
	seedsRange SeedsRange,
	seedsBySoils *shared.BinarySearchTree[Map],
	soilsByFertilizers *shared.BinarySearchTree[Map],
	fertilizersByWaters *shared.BinarySearchTree[Map],
	watersByLights *shared.BinarySearchTree[Map],
	lightsByTemperatures *shared.BinarySearchTree[Map],
	temperaturesByHumidities *shared.BinarySearchTree[Map],
	humiditiesByLocations *shared.BinarySearchTree[Map],
) int {
	soils := []int{}
	for i := range shared.Times(seedsRange.End - seedsRange.Start + 1) {
		r := FindLowestMappedValue(seedsRange.Start+i, seedsBySoils)
		soils = append(soils, r)
	}

	fertilizers := []int{}
	for _, soil := range soils {
		r := FindLowestMappedValue(soil, soilsByFertilizers)
		fertilizers = append(fertilizers, r)
	}

	waters := []int{}
	for _, fertilizer := range fertilizers {
		r := FindLowestMappedValue(fertilizer, fertilizersByWaters)
		waters = append(waters, r)
	}

	lights := []int{}
	for _, water := range waters {
		r := FindLowestMappedValue(water, watersByLights)
		lights = append(lights, r)

	}

	temperatures := []int{}
	for _, light := range lights {
		r := FindLowestMappedValue(light, lightsByTemperatures)
		temperatures = append(temperatures, r)
	}

	humidities := []int{}
	for _, temperature := range temperatures {
		r := FindLowestMappedValue(temperature, temperaturesByHumidities)
		humidities = append(humidities, r)
	}

	locations := []int{}
	for _, humidity := range humidities {
		r := FindLowestMappedValue(humidity, humiditiesByLocations)
		locations = append(locations, r)
	}

	return shared.FindLowestInteger(locations)
}
