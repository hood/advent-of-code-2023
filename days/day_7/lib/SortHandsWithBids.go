package lib

import "sort"

type HandsSorter []HandWithBid

func (a HandsSorter) Len() int      { return len(a) }
func (a HandsSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a HandsSorter) Less(i, j int) bool {
	if a[i].Hand.NumberScore() != a[j].Hand.NumberScore() {
		return a[i].Hand.NumberScore() < a[j].Hand.NumberScore()
	}

	return a[i].Hand.Stringified() < a[j].Hand.Stringified()
}

func SortHandsWithBids(hands []HandWithBid) {
	sort.Sort(HandsSorter(hands))
}
