package lib

import "strings"

type Node struct {
	ID    string
	Left  *Node
	Right *Node
}

func (n *Node) AddLeft(ID string) {
	n.Left = &Node{
		ID: ID,
	}
}

func (n *Node) AddRight(ID string) {
	n.Right = &Node{
		ID: ID,
	}
}

func Bfs(head *Node, target string) (bool, int, *Node) {
	current := head

	if current == nil || current.ID == "" {
		return false, -1, nil
	}

	toVisitQueue := make([]*Node, 0)
	visitedList := make([]*Node, 0)

	toVisitQueue = append(toVisitQueue, head)

	level := 0

	for len(toVisitQueue) > 0 {
		level++

		// Dequeue.
		n := toVisitQueue[0]
		toVisitQueue = toVisitQueue[1:]

		// Mark as visited.
		visitedList = append(visitedList, n)

		if n.ID == target {
			return true, level, n
		}

		if n.Left != nil {
			toVisitQueue = append(toVisitQueue, n.Left)
		}

		if n.Right != nil {
			toVisitQueue = append(toVisitQueue, n.Right)
		}
	}

	return true, level, current
}

type Map struct {
	Head *Node
}

func MapFromLines(lines []string) *Map {
	myMap := &Map{}

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

		found, _, foundnode := Bfs(myMap.Head, nodeID)
		if found {
			foundnode.AddLeft(left)
			foundnode.AddRight(right)
		} else {
			node := &Node{
				ID: nodeID,
			}

			node.AddLeft(left)
			node.AddRight(right)

			if myMap.Head == nil {
				myMap.Head = node
			}
		}

	}

	return myMap
}
