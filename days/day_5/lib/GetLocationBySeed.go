package lib

func GetLocationBySeed(
	seed int,
	seedsBySoils []Map,
	soilsByFertilizers []Map,
	fertilizersByWaters []Map,
	watersByLights []Map,
	lightsByTemperatures []Map,
	temperaturesByHumidities []Map,
	humiditiesByLocations []Map,
) int {
	soil := FindLowestMappedValue(seed, seedsBySoils)
	fertilizer := FindLowestMappedValue(soil, soilsByFertilizers)
	water := FindLowestMappedValue(fertilizer, fertilizersByWaters)
	light := FindLowestMappedValue(water, watersByLights)
	temperature := FindLowestMappedValue(light, lightsByTemperatures)
	humidity := FindLowestMappedValue(temperature, temperaturesByHumidities)
	location := FindLowestMappedValue(humidity, humiditiesByLocations)

	println(
		"\n\nseed", seed,
		"\nsoil", soil,
		"\nfertilizer", fertilizer,
		"\nwater", water,
		"\nlight", light,
		"\ntemperature", temperature,
		"\nhumidity", humidity,
		"\nlocation", location,
	)

	return location
}
