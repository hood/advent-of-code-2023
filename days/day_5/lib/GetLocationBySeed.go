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
