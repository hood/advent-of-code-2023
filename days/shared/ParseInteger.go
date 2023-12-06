package shared

import (
	"strconv"
	"strings"
)

func ParseInteger(number string) int {
	s, err := strconv.Atoi(strings.TrimSpace(number))
	if err != nil {
		panic(err)
	}

	return s
}
