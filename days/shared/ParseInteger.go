package shared

import (
	"strconv"
)

func ParseInteger(number string) int {
	s, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}

	return s
}
