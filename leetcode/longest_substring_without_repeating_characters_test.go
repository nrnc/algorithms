package leetcode

import "testing"

type Problem3 struct {
	input  string
	output int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []Problem3{
		{
			input:  "abcabcbb",
			output: 3,
		},
		{
			input:  "bbbbb",
			output: 1,
		},
		{
			input:  "pwwkew",
			output: 3,
		},
		{
			input:  "babf",
			output: 3,
		},
	}
	for _, v := range testCases {
		ans := LengthOfLongestSubstring(v.input)
		if ans != v.output {
			t.Errorf("expected %d got %d", v.output, ans)
		}
	}
}
