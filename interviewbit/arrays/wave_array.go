package arrays

import "sort"

type IntArray []int

func (r IntArray) Less(i, j int) bool {
	return r[i] < r[j]
}
func (r IntArray) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r IntArray) Len() int {
	return len(r)
}

func WaveArray(A []int) []int {
	n := len(A)
	sort.Sort(IntArray(A))
	for i := 0; i < n-1; i += 2 {
		A[i], A[i+1] = A[i+1], A[i]
	}

	return A
}
