package binarysearch

import "sort"

func SmallerOrEqualElements(A []int, B int) int {
	i := sort.Search(len(A), func(i int) bool { return A[i] > B })
	return i
}
