package lib

import "adventofcode2023/days/shared"

func FindLowestMappedValue(value int, mappings *shared.BinarySearchTree[Map]) int {
	result := -1

	found, foundValue := mappings.Find(value, func(node *shared.BinarySearchTreeNode[Map]) bool {
		return node.Value.FindMapping(value) != -1
	})

	if found {
		result = foundValue.FindMapping(value)
	}

	if result == -1 {
		return value
	}

	return result
}
