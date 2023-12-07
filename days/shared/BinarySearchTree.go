package shared

type BinarySearchTreeNode[T any] struct {
	Value      T
	Left       *BinarySearchTreeNode[T]
	Right      *BinarySearchTreeNode[T]
	LessThan   func(other interface{}) bool
	BiggerThan func(other interface{}) bool
	Equals     func(other interface{}) bool
}

type BinarySearchTree[T any] struct {
	Root *BinarySearchTreeNode[T]
}

func NewBinarySearchTree[T any]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (node *BinarySearchTreeNode[T]) Insert(value T) {
	if node.BiggerThan(value) {
		if node.Left == nil {
			node.Left = &BinarySearchTreeNode[T]{Value: value}
		} else {
			node.Left.Insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &BinarySearchTreeNode[T]{Value: value}
		} else {
			node.Right.Insert(value)
		}
	}
}

func (node *BinarySearchTreeNode[T]) Search(value T) *BinarySearchTreeNode[T] {
	if node.Equals(value) {
		return node
	}

	if node.BiggerThan(value) {
		if node.Left == nil {
			return nil
		}

		return node.Left.Search(value)
	}

	if node.Right == nil {
		return nil
	}

	return node.Right.Search(value)
}

func (tree *BinarySearchTree[T]) Insert(value T) {
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode[T]{Value: value}
		return
	}

	tree.Root.Insert(value)
}

func (tree *BinarySearchTree[T]) Search(value T) *BinarySearchTreeNode[T] {
	if tree.Root == nil {
		return nil
	}

	return tree.Root.Search(value)
}

// func (tree *BinarySearchTree[T]) Find(node T, compare func(nodeValue T, value T) bool) (bool, T) {
// 	var defaultValue T

// 	if tree.Root == nil {
// 		return false, defaultValue
// 	}

// 	currentNode := tree.Root
// 	for currentNode != nil {
// 		if compare(currentNode.Value, node) {
// 			return true, currentNode.Value
// 		} else if compare(node, currentNode.Value) {
// 			currentNode = currentNode.Left
// 		} else {
// 			currentNode = currentNode.Right
// 		}
// 	}

// 	return false, defaultValue
// }

func (tree *BinarySearchTree[T]) Find(value interface{}, compare func(nodeValue T, value T) bool) (bool, T) {
	var defaultValue T

	if tree.Root == nil {
		return false, defaultValue
	}

	currentNode := tree.Root
	for currentNode != nil {
		if compare(currentNode.Value, value) {
			return true, currentNode.Value
		} else if compare(value, currentNode.Value) {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}

	return false, defaultValue
}
