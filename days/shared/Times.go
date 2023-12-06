package shared

func Times(number int) []int {
	times := []int{}

	for i := 0; i < number; i++ {
		times = append(times, i)
	}

	return times
}
