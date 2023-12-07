package shared

import "time"

func PartRunner(day int, part int, fn func()) {
	println("\n\n***** Day", day, ".", part, "****")

	start := time.Now()

	fn()

	elapsed := time.Since(start)

	println("(", "Execution time:", elapsed, ")")
}
