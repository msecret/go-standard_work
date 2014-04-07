package work

type comparsion func(a, b float64) bool

type ComparableFloat64 interface {
	Compare(a ComparatorFloat64) int
}

type ComparatorFloat64 struct {
	Val float64
}

func (c ComparatorFloat64) Compare(a ComparatorFloat64) int {
	if c.Val > a.Val {
		return 1
	} else if c.Val < a.Val {
		return -1
	} else {
		return 0
	}
}

func SelectionSortFloat64(list []ComparatorFloat64) []ComparatorFloat64 {
	var rlist []ComparatorFloat64

	i := 0
	j := i + 1
	rlist = list
	for i < len(rlist) {
		j = i + 1
		Val := rlist[i]

		for ; j < len(rlist); j++ {
			if rlist[j].Compare(Val) == -1 {
				temp := rlist[i]
				rlist[i] = rlist[j]
				rlist[j] = temp
				Val = rlist[i]
			}
		}
		i += 1
	}

	return rlist
}

func MergeSortFloat64(list []ComparatorFloat64) []ComparatorFloat64 {
	if len(list) < 2 {
		return list[:]
	} else {
		mid := len(list) / 2
		left := MergeSortFloat64(list[:mid])
		right := MergeSortFloat64(list[mid:])

		return merge(left, right)
	}
}

func merge(lA, lB []ComparatorFloat64) []ComparatorFloat64 {
	return lA
}
