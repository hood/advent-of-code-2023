package shared

import (
	"fmt"
	"time"

	"golang.design/x/clipboard"
)

func RunSolution(solution func(callback func(r interface{}))) {
	start := time.Now()

	solution(func(result interface{}) {
		elapsed := time.Since(start)

		fmt.Println(fmt.Sprintf("Result -> %v (took %vms)", result, elapsed.Milliseconds()))

		clipboard.Init()
		clipboard.Write(clipboard.FmtText, []byte(fmt.Sprintf("%v", result)))
	})
}
