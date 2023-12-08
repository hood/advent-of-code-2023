package lib

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

	if current.ID == "" {
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
