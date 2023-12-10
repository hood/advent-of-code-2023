package lib

import "math"

// import (
// 	"math"
// )

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

// func Lagrange(nums []int) int {
// 	n := len(nums)
// 	res := 0

// 	for i, x := range nums {
// 		res += x * comb(n, i) * int(math.Pow(-1, float64(n-1-i)))
// 	}

// 	return res
// }

// func Lagrange(nums []int) int {
// 	n := len(nums)
// 	res := 0

// 	for i, y := range nums {
// 		lx := 1
// 		for j, x := range nums {
// 			if i != j {
// 				lx *= ((n + 1) - x) / (nums[i] - x)
// 			}
// 		}
// 		res += y * lx
// 	}

// 	return res
// }

func Lagrange(nums []int) int {
	n := len(nums)
	res := 0

	for i, num := range nums {
		res += num * comb(n, i) * int(math.Pow(-1, float64(n-1-i))) // **(n - 1 - i)
	}

	return res
}

// def Lagrange1(nums):
//     n=len(nums)
//     res=0
//     for i,x in enumerate(nums):
//         res+=x*comb(n,i)*(-1)**(n-1-i)
//     return res
