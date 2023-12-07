package shared

import (
	"fmt"
	"time"
)

func RunSolution(solution func(callback func(r interface{}))) {
	start := time.Now()

	solution(func(result interface{}) {
		elapsed := time.Since(start)

		fmt.Println(fmt.Sprintf("Result -> %v (took %vms)", result, elapsed.Milliseconds()))
	})
}
