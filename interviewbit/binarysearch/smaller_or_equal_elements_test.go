package binarysearch

import "testing"

type ProblemSmallerOrEqualElements struct {
	input  []int
	output int
}

func TestSmallerOrEqualElements(t *testing.T) {
	testCases := []ProblemSmallerOrEqualElements{
		{
			input:  []int{1, 3, 4, 4, 6},
			output: 4,
		},
		{
			input:  []int{1, 2, 5, 5},
			output: 2,
		},
	}
	ans1 := SmallerOrEqualElements(testCases[0].input, 4)
	ans2 := SmallerOrEqualElements(testCases[1].input, 3)
	if ans1 != testCases[0].output || ans2 != testCases[1].output {
		t.Error("test case Smaller or equal elements is failed")
	}
}
