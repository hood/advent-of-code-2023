package main

import (
	"adventofcode2023/days/day_1_1"
	"adventofcode2023/days/day_1_2"
	"adventofcode2023/days/day_2"
	"adventofcode2023/days/day_3"
	"adventofcode2023/days/day_4"
	"adventofcode2023/days/day_5"
)

func main() {
	println("Starting...")

	day_1_1.Entry()
	day_1_2.Entry()

	day_2.Day2()

	day_3.Day3()

	day_4.Day4()

	day_5.Day5()

	println("\n\nDone!")
}
