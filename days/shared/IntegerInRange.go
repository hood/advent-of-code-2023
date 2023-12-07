package shared

func IntegersInRange(start int, end int) []int {
	integers := []int{}

	for i := start; i <= end; i++ {
		integers = append(integers, i)
	}

	return integers
}
