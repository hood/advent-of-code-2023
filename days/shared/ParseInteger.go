package shared

import (
	"strconv"
	"strings"
)

func ParseInteger(number string) int {
	if number == "" {
		panic("stringified number is empty")
	}

	s, err := strconv.Atoi(strings.TrimSpace(number))
	if err != nil {
		panic(err)
	}

	return s
}
