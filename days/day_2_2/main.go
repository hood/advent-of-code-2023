package day_2_2

import "os"

// Day2.2 entry point.
func Entry() {
	println("\n\n***** Day 2.2 ****")

	fileContent, error := os.ReadFile("./days/day_2_2/input.txt")
	if error != nil {
		panic(error)
	}

	println("File content:")
}
