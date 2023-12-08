package lib

import "strings"

type Node struct {
	ID    string
	Left  *Node
	Right *Node
}

func (n *Node) AddLeft(node *Node) {
	n.Left = node
}

func (n *Node) AddRight(node *Node) {
	n.Right = node
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
	nodes := make(map[string]*Node)

	// Build nodes.
	for _, line := range lines {
		if line == "" {
			continue
		}

		nodeID := strings.Split(line, " = ")[0]

		_, existingNodeFound := nodes[nodeID]
		if existingNodeFound == false {
			nodes[nodeID] = &Node{
				ID: nodeID,
			}

			if myMap.Head == nil {
				myMap.Head = nodes[nodeID]
			}
		}
	}

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

		_, nodeFound := nodes[nodeID]
		leftNode, leftNodeFound := nodes[left]
		rightNode, rightNodeFound := nodes[right]

		if !nodeFound || !leftNodeFound || !rightNodeFound {
			panic("Node(s) not found, shouldn't happen!")
		}

		nodes[nodeID].AddLeft(leftNode)
		nodes[nodeID].AddRight(rightNode)
	}

	return myMap
}
