package lib

import (
	"adventofcode2023/days/shared"
	"testing"
)

func TestMap(t *testing.T) {
	aaa := Node{
		ID: "AAA",
	}

	aaa.AddLeft("BBB")
	aaa.AddRight("CCC")

	println(
		aaa.Left.ID,
		aaa.Right.ID,
	)

	bbb := aaa.Left
	bbb.AddLeft("DDD")
	bbb.AddRight("EEE")

	ccc := aaa.Right
	ccc.AddLeft("ZZZ")
	ccc.AddRight("GGG")

	ddd := bbb.Left
	ddd.AddLeft("DDD")
	ddd.AddRight("DDD")

	eee := bbb.Right
	eee.AddLeft("EEE")
	eee.AddRight("EEE")

	ggg := ccc.Left
	ggg.AddLeft("GGG")
	ggg.AddRight("GGG")

	zzz := ccc.Right
	zzz.AddLeft("ZZZ")
	zzz.AddRight("ZZZ")

	found, _, foundNode := Bfs(&aaa, "ZZZ")

	shared.AssertEqual(t, true, found)
	shared.AssertEqual(t, "ZZZ", foundNode.ID)
}
