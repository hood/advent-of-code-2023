package lib

func GetLocationBySeed(
	seed int,
	humiditiesByLocations []Map,
	temperaturesByHumidities []Map,
	lightsByTemperatures []Map,
	watersByLights []Map,
	fertilizersByWaters []Map,
	soilsByFertilizers []Map,
	seedsBySoils []Map,
) int {
	soil := FindMappedValue(seed, seedsBySoils)
	fertilizer := FindMappedValue(soil, soilsByFertilizers)
	water := FindMappedValue(fertilizer, fertilizersByWaters)
	light := FindMappedValue(water, watersByLights)
	temperature := FindMappedValue(light, lightsByTemperatures)
	humidity := FindMappedValue(temperature, temperaturesByHumidities)
	location := FindMappedValue(humidity, humiditiesByLocations)

	return location
}
