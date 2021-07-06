package leetcode

import "testing"

type Problem7 struct {
	input  int
	output int
}

func TestReverse(t *testing.T) {
	testCases := []Problem7{
		{
			input:  123,
			output: 321,
		},
		{
			input:  -123,
			output: -321,
		},
	}
	for _, test := range testCases {
		ans := Reverse(test.input)
		if ans != test.output {
			t.Errorf("expected %d got %d", test.output, ans)
		}
	}
}
