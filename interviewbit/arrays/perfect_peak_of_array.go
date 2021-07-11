package arrays

func PerfectPeak(A []int) int {
	n := len(A)
	maxArrayFromLeft := make([]int, n)
	minArrayFromRight := make([]int, n)
	maxArrayFromLeft[0] = A[0]
	minArrayFromRight[n-1] = A[n-1]

	for i := 1; i < n; i++ {
		if maxArrayFromLeft[i-1] > A[i] {
			maxArrayFromLeft[i] = maxArrayFromLeft[i-1]
		} else {
			maxArrayFromLeft[i] = A[i]
		}
	}
	for i := n - 2; i >= 0; i-- {
		if minArrayFromRight[i+1] < A[i] {
			minArrayFromRight[i] = minArrayFromRight[i+1]
		} else {
			minArrayFromRight[i] = A[i]
		}
	}

	for i := 1; i < n-1; i++ {
		if maxArrayFromLeft[i-1] != maxArrayFromLeft[i] && minArrayFromRight[i] != minArrayFromRight[i+1] {
			return 1
		}
	}

	return 0
}
