package leetcode

import "testing"

type Input4 struct {
	nums1 []int
	nums2 []int
}

type Problem4 struct {
	input  Input4
	output float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	testCases := []Problem4{
		{
			input: Input4{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			output: 2.00000,
		},
		{
			input: Input4{
				nums1: []int{1, 3},
				nums2: []int{2, 4},
			},
			output: 2.50000,
		},
		{
			input: Input4{
				nums1: []int{0, 0},
				nums2: []int{0, 0},
			},
			output: 0.00000,
		},
		{
			input: Input4{
				nums1: []int{2},
				nums2: []int{},
			},
			output: 2.00000,
		},
	}

	for _, test := range testCases {
		output := FindMedianSortedArrays(test.input.nums1, test.input.nums2)
		if output != test.output {
			t.Errorf("expected %f got %f", test.output, output)
		}
	}
}
