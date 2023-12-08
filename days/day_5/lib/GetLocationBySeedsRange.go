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
	mappings := make([]int, 0)

	stepSize := 1000000
	i := 0

	// Try to iterate over the points with a large step size to avoid calculating
	// obvious mappings. If the mappings start to differ, reduce the step size
	// and try again.
	for i*stepSize < len(points) {
		if i > 1 &&
			mappings[i-1]-mappings[i-2] != stepSize &&
			stepSize > 1 {
			if stepSize < 1 {
				panic("WTF")
			}

			stepSize = stepSize / 10

			continue
		}

		mappings = append(mappings, FindLowestMappedValue(points[i*stepSize], maps))

		i++
	}

	// for _, point := range points {
	// 	r := FindLowestMappedValue(point, maps)
	// 	mappings = append(mappings, r)
	// }

	return mappings
}
