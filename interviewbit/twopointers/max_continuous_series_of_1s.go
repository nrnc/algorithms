package twopointers

func MaxContinuousSeriesOf1s(A []int, B int) []int {
	start := 0
	end := 0
	zeroCount := 0
	length := 0
	low := 0
	high := 0
	for ; end < len(A); end++ {
		if A[end] == 0 {
			zeroCount++
		}
		for zeroCount > B {
			if A[start] == 0 {
				zeroCount--
			}
			start++
		}
		if end-start+1 > length {
			length = end - start + 1
			low = start
			high = end
		}
	}
	ans := make([]int, 0)
	for i := low; i <= high; i++ {
		ans = append(ans, i)
	}

	return ans
}
