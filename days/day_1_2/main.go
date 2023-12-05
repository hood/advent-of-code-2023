package day_1_2

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

var digits = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var stringifiedDigits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// Day1.2 entry point.
func Entry() {
	println("\n\n***** Day 2 ****")

	fileContent, error := os.ReadFile("./days/day_1_2/input.txt")
	if error != nil {
		panic(error)
	}

	allValidDigits := append(stringifiedDigits, digits...)

	lines := strings.Split(string(fileContent), "\n")

	calibrationValues := make([]string, len(lines))

	for lineIndex, line := range lines {
		firstIndex, lastIndex := len(line), -1
		first, last := "", ""

		for _, validDigit := range allValidDigits {
			startIndex := strings.Index(line, validDigit)
			endIndex := strings.LastIndex(line, validDigit)

			if startIndex == -1 {
				continue
			}

			// If it's in a place before the earliest digit, register it as first.
			if startIndex < firstIndex {
				firstIndex = startIndex
				first = validDigit
			}

			// If it's in a place after the latest digit, register it as last.
			if startIndex > lastIndex {
				lastIndex = startIndex
				last = validDigit
			}

			// If it's in a place before the earliest digit, register it as first.
			if endIndex < firstIndex {
				firstIndex = endIndex
				first = validDigit
			}

			// If it's in a place after the latest digit, register it as last.
			if endIndex > lastIndex {
				lastIndex = endIndex
				last = validDigit
			}
		}

		calibrationValues[lineIndex] = digitToStringifiedInteger(first) + digitToStringifiedInteger(last)
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

func digitToStringifiedInteger(digit string) string {
	if slices.Contains(digits, digit) {
		return digit
	}

	switch digit {
	case "one":
		return "1"

	case "two":
		return "2"

	case "three":
		return "3"

	case "four":
		return "4"

	case "five":
		return "5"

	case "six":
		return "6"

	case "seven":
		return "7"

	case "eight":
		return "8"

	case "nine":
		return "9"

	default:
		panic("Shouldn't ever get here! Value is: " + digit)
	}
}
