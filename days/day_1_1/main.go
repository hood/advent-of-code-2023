package day_1_1

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Day1.1 entry point.
func Entry() {
	println("\n\n***** Day 1 ****")

	fileContent, error := os.ReadFile("./days/day_1_1/input.txt")
	if error != nil {
		panic(error)
	}

	lines := strings.Split(string(fileContent), "\n")

	calibrationValues := make([]string, len(lines))

	for lineIndex, line := range lines {
		lineCalibrationValues := make([]string, 2)

		for _, char := range line {
			parsedDigit := unicode.IsDigit(char)

			if parsedDigit {
				if lineCalibrationValues[0] == "" {
					lineCalibrationValues[0] = string(char)
				}

				lineCalibrationValues[1] = string(char)
			}
		}

		calibrationValues[lineIndex] = lineCalibrationValues[0] + lineCalibrationValues[1]
	}

	result := 0

	for _, calibrationValue := range calibrationValues {
		parsedValue, error := strconv.Atoi(calibrationValue)
		if error != nil {
			panic(error)
		}

		result += parsedValue
	}

	println("Result", "->", result)
}
