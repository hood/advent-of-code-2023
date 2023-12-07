package lib

import "sort"

type HandsSorter []HandWithBid

func (a HandsSorter) Len() int      { return len(a) }
func (a HandsSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a HandsSorter) Less(i, j int) bool {
	if a[i].Hand.Score() != a[j].Hand.Score() {
		return a[i].Hand.Score() < a[j].Hand.Score()
	}

	return a[i].Hand.Stringified > a[j].Hand.Stringified
}

func SortHandsWithBids(hands []HandWithBid) {
	sort.Sort(HandsSorter(hands))
}

// JJJJJ 0
// TTTTJ 1
// TTTT8 1
// TTTAT 1
// TTT7T 1
// TTT2T 1
// TQTTT 1
// TKTTT 1
// TKKKK 1
// T5555 1
// QTQQQ 1
// QQQQA 1
// QQQJQ 1
// QQQ7Q 1
// QQ8QQ 1
// Q8888 1
// Q6666 1
// KQQQQ 1
// KKKK2 1
// KKK8K 1
// KKK5K 1
// KJKKK 1
// J8JJJ 1
// J8888 1
// J7777 1
// AJAAA 1
// AATAA 1
// AAQAA 1
// AAAKA 1
// AAA8A 1
// AA9AA 1
// A5AAA 1
// A3AAA 1
// 9A999 1
// 99T99 1
// 99J99 1
// 99992 1
// 99799 1
// 98999 1
// 96999 1
// 95999 1
// 93333 1
