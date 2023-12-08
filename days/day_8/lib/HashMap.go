package lib

import "strings"

type HashMapNode struct {
	L string
	R string
}

type HashMap map[string]HashMapNode

func HashMapFromLines(lines []string) HashMap {
	hashMap := make(HashMap)

	// Build edges.
	for _, line := range lines {
		if line == "" {
			continue
		}

		nodeID_leftRight := strings.Split(line, " = ")
		nodeID := nodeID_leftRight[0]
		leftRight := nodeID_leftRight[1]

		left_right := strings.Split(leftRight, ", ")
		left := left_right[0]
		right := left_right[1]
		left = strings.Split(left, "(")[1]
		right = strings.Split(right, ")")[0]

		hashMap[nodeID] = HashMapNode{
			L: left,
			R: right,
		}
	}

	return hashMap
}
