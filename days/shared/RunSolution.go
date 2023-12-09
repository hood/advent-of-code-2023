package shared

import (
	"fmt"
	"time"

	"golang.design/x/clipboard"
)

func RunSolution(
	day int,
	part int,
	solution func(callback func(r interface{})),
) {
	fmt.Printf("\n\n**** Day %v.%b ****", day, part)

	start := time.Now()

	solution(func(result interface{}) {
		elapsed := time.Since(start)

		fmt.Println(fmt.Sprintf("Result -> %v (took %v)", result, elapsed.String()))

		clipboard.Init()
		clipboard.Write(clipboard.FmtText, []byte(fmt.Sprintf("%v", result)))

		recordTime := GetRecord(day, part)

		if elapsed.Microseconds() < recordTime || recordTime == 0 {
			RegisterRecord(day, part, elapsed.Microseconds())
		}
	})
}
