package lib

import (
	"math"
)

// func comb(n, k int) int {
// 	f := float64(n)

// 	for i := 1; i < k; i++ {
// 		f *= float64(n-i) / float64(i+1)
// 	}

// 	return int(f + 0.5)
// }

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func comb(n, k int) int {
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func Lagrange1(nums []int) int {
	n := len(nums)
	res := 0

	for i, x := range nums {
		res += x * comb(n, i) * int(math.Pow(-1, float64(n-1-i)))
	}

	return res
}

func Lagrange2(nums []int) float64 {
	n := len(nums)
	res := 0.0

	for i, x := range nums {
		res += float64(x) * float64(comb(n, i+1)) * math.Pow(-1, float64(i))
	}

	return res
}
