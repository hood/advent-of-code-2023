package lib

import "testing"

func TestMap(t *testing.T) {
	aaa := Node{
		ID: "AAA",
	}

	aaa.AddLeft(Node{
		ID: "BBB",
	})
	aaa.AddRight(Node{
		ID: "CCC",
	})

	bbb := aaa.Left
	bbb.AddLeft(Node{
		ID: "DDD",
	})
	bbb.AddRight(Node{
		ID: "EEE",
	})

	ccc := aaa.Right
	ccc.AddLeft(Node{
		ID: "ZZZ",
	})

	ddd := bbb.Left
	ddd.AddLeft(Node{
		ID: "DDD",
	})
	ddd.AddRight(Node{
		ID: "DDD",
	})

	eee := bbb.Right
	eee.AddLeft(Node{
		ID: "EEE",
	})
	eee.AddRight(Node{
		ID: "EEE",
	})

	ggg := ccc.Left
	ggg.AddLeft(Node{
		ID: "GGG",
	})
	ggg.AddRight(Node{
		ID: "GGG",
	})

	zzz := ccc.Right
	zzz.AddLeft(Node{
		ID: "ZZZ",
	})
	zzz.AddRight(Node{
		ID: "ZZZ",
	})

	result := Bfs(&aaa, "ZZZ")

	if result.ID != "ZZZ" {
		t.Errorf("Expected ZZZ, got %s", result.ID)
	}
}

// AAA = (BBB, CCC)
// BBB = (DDD, EEE)
// CCC = (ZZZ, GGG)
// DDD = (DDD, DDD)
// EEE = (EEE, EEE)
// GGG = (GGG, GGG)
// ZZZ = (ZZZ, ZZZ)

// func findById(root *Node, id string) *Node {
// 	queue := make([]*Node, 0)
// 	queue = append(queue, root)
// 	for len(queue) > 0 {
// 					nextUp := queue[0]
// 					queue = queue[1:]
// 					if nextUp.id == id {
// 									return nextUp
// 					}
// 					if len(nextUp.children) > 0 {
// 									for _, child := range nextUp.children {
// 													queue = append(queue, child)
// 									}
// 					}
// 	}
// 	return nil
// }
