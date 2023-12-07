package shared

import (
	"os"
	"strings"
)

func ReadFile(path string) []string {
	fileContent, error := os.ReadFile(path)
	if error != nil {
		panic(error)
	}

	lines := strings.Split(string(fileContent), "\n")

	return lines
}
