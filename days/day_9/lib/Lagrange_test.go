package lib

import (
	"adventofcode2023/days/shared"
	"testing"
)

func TestLagrange(t *testing.T) {
	values := []int{0, 1, 1, 0}

	shared.AssertEqual(t, -2, Lagrange(values))
}
