package leetcode

import (
	"testing"
)

type input struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

type output []int

type problem struct {
	input  input
	output output
}

func slicesEqual(s1 []int, s2 []int) bool {
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}
func TestMergeSortedArray(t *testing.T) {
	testCases := []problem{
		{
			input: input{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			output: []int{1, 2, 2, 3, 5, 6},
		},
		{
			input: input{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			output: []int{1},
		},
		{
			input: input{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			output: []int{1},
		},
	}
	for _, test := range testCases {
		MergeSortedArray(test.input.nums1, test.input.m, test.input.nums2, test.input.n)
		if !slicesEqual(test.input.nums1, test.output) {
			t.Errorf("expected %v but got %v", test.output, test.input.nums1)
		}
	}
}
