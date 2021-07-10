package leetcode

import "testing"

type Problem14 struct {
	input  []string
	output string
}

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []Problem14{
		{
			input:  []string{"flower", "flow", "flight"},
			output: "fl",
		},
		{
			input:  []string{"dog", "racecar", "car"},
			output: "",
		},
		{
			input:  []string{"ab", "a"},
			output: "a",
		},
	}

	for _, test := range testCases {
		ans := LongestCommonPrefix(test.input)

		if ans != test.output {
			t.Errorf("expected %s but got %s", test.output, ans)
		}
	}
}
