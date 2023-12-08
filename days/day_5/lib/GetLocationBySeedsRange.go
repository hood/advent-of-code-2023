package lib

import (
	"adventofcode2023/days/shared"
)

func GetLocationBySeedsRange(
	seedsRange SeedsRange,
	seedsBySoils []Map,
	soilsByFertilizers []Map,
	fertilizersByWaters []Map,
	watersByLights []Map,
	lightsByTemperatures []Map,
	temperaturesByHumidities []Map,
	humiditiesByLocations []Map,
	lowestSeedLocation int,
) int {
	soils := []int{}
	for i := range shared.Times(seedsRange.End - seedsRange.Start + 1) {
		r := FindLowestMappedValue(seedsRange.Start+i, seedsBySoils)
		soils = append(soils, r)
	}

	fertilizers := getMappings(soils, soilsByFertilizers)
	waters := getMappings(fertilizers, fertilizersByWaters)
	lights := getMappings(waters, watersByLights)
	temperatures := getMappings(lights, lightsByTemperatures)
	humidities := getMappings(temperatures, temperaturesByHumidities)

	if lowestSeedLocation != -1 &&
		lowestSeedLocation < humiditiesByLocations[0].SourceRangeStart {
		return lowestSeedLocation
	}

	locations := getMappings(humidities, humiditiesByLocations)

	return shared.FindLowestInteger(locations)
}

func getMappings(points []int, maps []Map) []int {
	mappings := []int{}

	for _, point := range points {
		r := FindLowestMappedValue(point, maps)
		mappings = append(mappings, r)
	}

	return mappings
}
