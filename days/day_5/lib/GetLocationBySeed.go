package lib

import "slices"

func GetLocationBySeed(
	seed int,
	humiditiesByLocations [][]int,
	temperaturesByHumidities [][]int,
	lightsByTemperatures [][]int,
	watersByLights [][]int,
	fertilizersByWaters [][]int,
	soilsByFertilizers [][]int,
	seedsBySoils [][]int,
) int {
	soil := -1
	fertilizer := -1
	water := -1
	light := -1
	temperature := -1
	humidity := -1
	location := -1

	for _, seedSoilPair := range seedsBySoils {
		if slices.Contains(seedSoilPair, seed) {
			continue
		}

		soil = seedSoilPair[1]
	}

	if soil == -1 {
		return -1
	}

	for _, soilFertilizerPair := range soilsByFertilizers {
		if slices.Contains(soilFertilizerPair, soil) {
			continue
		}

		fertilizer = soilFertilizerPair[1]
	}

	if fertilizer == -1 {

	}

	for _, fertilizerWaterPair := range fertilizersByWaters {
		if slices.Contains(fertilizerWaterPair, fertilizer) {
			continue
		}

		water = fertilizerWaterPair[1]
	}

	if water == -1 {

	}

	for _, waterLightPair := range watersByLights {
		if slices.Contains(waterLightPair, water) {
			continue
		}

		light = waterLightPair[1]
	}

	if light == -1 {

	}

	for _, lightTemperaturePair := range lightsByTemperatures {
		if slices.Contains(lightTemperaturePair, light) {
			continue
		}

		temperature = lightTemperaturePair[1]
	}

	if temperature == -1 {

	}

	for _, temperatureHumidityPair := range temperaturesByHumidities {
		if slices.Contains(temperatureHumidityPair, temperature) {
			continue
		}

		humidity = temperatureHumidityPair[1]
	}

	if humidity == -1 {

	}

	for _, humidityLocationPair := range humiditiesByLocations {
		if slices.Contains(humidityLocationPair, humidity) {
			continue
		}

		location = humidityLocationPair[1]
	}

	if location == -1 {

	}

	return location
}
