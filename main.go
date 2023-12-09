package main

import (
	"adventofcode2023/days/day_8"
	"adventofcode2023/days/day_9"
	"os"
)

func main() {
	println("Starting...")

	day := os.Args[1]
	part := os.Args[2]

	// day_1_1.Entry()
	// day_1_2.Entry()

	// day_2.Day2()

	// day_3.Day3()

	// day_4.Day4()

	// day_5.Day5()

	// day_6.Day6()

	// day_7.Day7()

	if day == "8" {
		if part == "1" {
			day_8.Day8Part1()
			return
		}

		if part == "2" {
			day_8.Day8Part2()
			return
		}
	}

	if day == "9" {
		if part == "1" {
			day_9.Day9Part1()
			return
		}

		// if part == "2" {
		// 	day_9.Day9Part2()
		// 	return
		// }
	}

	println("Nothing found!")
}
