package twopointers

import "sort"

type IntArray []int

func (ia IntArray) Len() int {
	return len(ia)
}
func (ia IntArray) Swap(i, j int) {
	ia[i], ia[j] = ia[j], ia[i]
}
func (ia IntArray) Less(i, j int) bool {
	return ia[i] < ia[j]
}
func PairWithGivenDifference(A []int, B int) int {
	sort.Sort(IntArray(A))
	i, j := 0, 1
	for i < len(A) && j < len(A) {
		if i != j && A[j]-A[i] == B {
			return 1
		} else if A[j]-A[i] < B {
			j++
		} else {
			i++
		}
	}
	return 0
}
