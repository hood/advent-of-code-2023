package lib

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
		lowestMappedSeed := seedSoilPair[0]

		if lowestMappedSeed > seed {
			break
		}

		if seedSoilPair[0] != seed {
			continue
		}

		soil = seedSoilPair[1]
	}

	if soil == -1 {
		soil = seed
	}

	for _, soilFertilizerPair := range soilsByFertilizers {
		lowestMappedSoil := soilFertilizerPair[0]

		if lowestMappedSoil > soil {
			break
		}

		if soilFertilizerPair[0] != soil {
			continue
		}

		fertilizer = soilFertilizerPair[1]
	}

	if fertilizer == -1 {
		fertilizer = soil
	}

	for _, fertilizerWaterPair := range fertilizersByWaters {
		lowestMappedFertilizer := fertilizerWaterPair[0]

		if lowestMappedFertilizer > fertilizer {
			break
		}

		if fertilizerWaterPair[0] != fertilizer {
			continue
		}

		water = fertilizerWaterPair[1]
	}

	if water == -1 {
		water = fertilizer
	}

	for _, waterLightPair := range watersByLights {
		lowestMappedWater := waterLightPair[0]

		if lowestMappedWater > water {
			break
		}

		if waterLightPair[0] != water {
			continue
		}

		light = waterLightPair[1]
	}

	if light == -1 {
		light = water
	}

	for _, lightTemperaturePair := range lightsByTemperatures {
		lowestMappedLight := lightTemperaturePair[0]

		if lowestMappedLight > light {
			break
		}

		if lightTemperaturePair[0] != light {
			continue
		}

		temperature = lightTemperaturePair[1]
	}

	if temperature == -1 {
		temperature = light
	}

	for _, temperatureHumidityPair := range temperaturesByHumidities {
		lowestMappedTemperature := temperatureHumidityPair[0]

		if lowestMappedTemperature > temperature {
			break
		}

		if temperatureHumidityPair[0] != temperature {
			continue
		}

		humidity = temperatureHumidityPair[1]
	}

	if humidity == -1 {
		humidity = temperature
	}

	for _, humidityLocationPair := range humiditiesByLocations {
		lowestMappedHumidity := humidityLocationPair[0]

		if lowestMappedHumidity > humidity {
			break
		}

		if humidityLocationPair[0] != humidity {
			continue
		}

		location = humidityLocationPair[1]
	}

	if location == -1 {
		location = humidity
	}

	return location
}
