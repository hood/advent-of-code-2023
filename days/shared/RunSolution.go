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

		fmt.Println(fmt.Sprintf("Result -> %v (took %v)", result, elapsed.String()))

		clipboard.Init()
		clipboard.Write(clipboard.FmtText, []byte(fmt.Sprintf("%v", result)))
	})

}
