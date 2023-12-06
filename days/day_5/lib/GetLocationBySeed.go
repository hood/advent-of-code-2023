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
	soil := FindMappedValue(seed, seedsBySoils)
	fertilizer := FindMappedValue(soil, soilsByFertilizers)
	water := FindMappedValue(fertilizer, fertilizersByWaters)
	light := FindMappedValue(water, watersByLights)
	temperature := FindMappedValue(light, lightsByTemperatures)
	humidity := FindMappedValue(temperature, temperaturesByHumidities)
	location := FindMappedValue(humidity, humiditiesByLocations)

	return location
}
