package lib

type Node struct {
	ID    string
	Left  *Node
	Right *Node
}

func (n *Node) AddLeft(node Node) {
	n.Left = &node
}

func (n *Node) AddRight(node Node) {
	n.Right = &node
}

func Bfs(head *Node, target string) *Node {
	// Find the node with the given id, then populate left and right.
	current := head

	if current.ID == "" {
		head = &Node{
			ID: target,
		}
	}

	toVisitQueue := []*Node{
		current,
	}
	visitedList := []*Node{}

	for len(toVisitQueue) > 0 {
		n := toVisitQueue[0]

		// Mark as visited.
		visitedList = append(visitedList, n)

		// Dequeue.
		toVisitQueue = toVisitQueue[1:]

		if n.Left != nil {
			toVisitQueue = append(toVisitQueue, n.Left)
		}

		if n.Right != nil {
			toVisitQueue = append(toVisitQueue, n.Right)
		}
	}

	return current
}

// type Map struct {
// 	Head Node
// }

// func NewMap(head Node) *Map {
// 	return &Map{
// 		Head: head,
// 	}
// }

// type NodeWithLevel struct {
// 	Node  *Node
// 	Level int
// }

// // https://faun.pub/2-different-ways-to-implement-bfs-in-golang-8399f5d2452d
// func (m *Map) Find(nodeID string) *Node {
// 	// Find the node with the given id, then populate left and right.
// 	current := m.Head

// 	if current.ID == "" {
// 		m.Head = Node{
// 			ID: nodeID,
// 		}
// 	}

// 	toVisitQueue := []NodeWithLevel{
// 		{
// 			Node:  &current,
// 			Level: 0,
// 		},
// 	}
// 	visitedList := []NodeWithLevel{}

// 	for len(toVisitQueue) > 0 {
// 		n := toVisitQueue[0]

// 		// Mark as visited.
// 		visitedList = append(visitedList, n)

// 		// Dequeue.
// 		toVisitQueue = toVisitQueue[1:]

// 		if n.Node.Left != nil {
// 			left := NodeWithLevel{
// 				Node:  n.Node.Left,
// 				Level: n.Level + 1,
// 			}

// 			toVisitQueue = append(toVisitQueue, left)

// 			right := NodeWithLevel{
// 				Node:  n.Node.Right,
// 				Level: n.Level + 1,
// 			}

// 			toVisitQueue = append(toVisitQueue, right)
// 		}
// 	}

// 	return &current
// }

// raw input looks like this.
// for example node AAA has left = BBB, right = CCC,
// node BBB has left = DDD, right = EEE, etc.
// AAA = (BBB, CCC)
// BBB = (DDD, EEE)
// CCC = (ZZZ, GGG)
// DDD = (DDD, DDD)
// EEE = (EEE, EEE)
// GGG = (GGG, GGG)
// ZZZ = (ZZZ, ZZZ)
