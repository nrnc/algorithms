package twopointers

func MaxOnesAfterModification(A []int, B int) int {
	start := 0
	end := 0
	zeroCount := 0
	ans := 0
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
		if end-start+1 > ans {
			ans = end - start + 1
		}
	}
	return ans
}
