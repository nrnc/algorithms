package twopointers

func IntersectionOfSortedArrays(A []int, B []int) []int {
	ans := make([]int, 0)
	var i, j int
	in, jn := len(A), len(B)
	for i < in && j < jn {
		if A[i] == B[j] {
			ans = append(ans, A[i])
			i++
			j++
		} else if A[i] < B[j] {
			i++
		} else {
			j++
		}
	}

	return ans
}
