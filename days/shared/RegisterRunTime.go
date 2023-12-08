package shared

import (
	"fmt"
	"os"
	"strings"
)

func RegisterRunTime(day int, part int, runTime int) {
	registeredLines := ReadFile("./run-times.txt")
	found := false

	for _, line := range registeredLines {
		if !strings.Contains(line, fmt.Sprintf("Day%v.%v", day, part)) {
			continue
		}

		if runTime < parseRunTime(line) {
			found = true

			break
		}
	}

	if !found {
		file, error := os.OpenFile("./run-times.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if error != nil {
			panic(error)
		}

		defer file.Close()

		_, error = file.WriteString(fmt.Sprintf("Day%v.%v, %vms\n", day, part, runTime))
		if error != nil {
			panic(error)
		}
	} else {
		fmt.Println(fmt.Sprintf("Day %v, part %v is not faster than the current record.", day, part))
	}
}

func parseRunTime(line string) int {
	time := strings.Split(line, ", ")[1]
	time = strings.ReplaceAll(time, "ms", "")
	time = strings.ReplaceAll(time, " ", "")

	intTime := 0
	fmt.Sscanf(time, "%d", &intTime)

	return intTime
}
