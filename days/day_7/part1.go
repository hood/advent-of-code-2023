package day_7

import (
	"adventofcode2023/days/day_7/lib"
	"adventofcode2023/days/shared"
	"sort"
)

func day7Part1() {
	println("\n\n***** Day 7.1 ****")

	lines := shared.ReadFile("./days/day_7/part_1_test_input.txt")

	handsWithBids := make([]lib.HandWithBid, len(lines))

	for i, line := range lines {
		handsWithBids[i] = lib.ParseHandWithBid(line)
	}

	sort.Sort(byscore(handsWithBids))

	for _, handWithBid := range handsWithBids {
		println(handWithBid.Hand.Stringified, handWithBid.Hand.Score(), handWithBid.Bid)
	}
}

type byscore []lib.HandWithBid

func (a byscore) Len() int      { return len(a) }
func (a byscore) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byscore) Less(i, j int) bool {
	if a[i].Hand.Score() != a[j].Hand.Score() {
		return a[i].Hand.Score() < a[j].Hand.Score()
	}

	return a[i].Hand.Stringified > a[j].Hand.Stringified
}
