package leetcode

import "testing"

type Input6 struct {
	s       string
	numRows int
}
type Problem6 struct {
	input  Input6
	output string
}

func TestConvertZigZag(t *testing.T) {
	testCases := []Problem6{
		{
			input:  Input6{s: "PAYPALISHIRING", numRows: 3},
			output: "PAHNAPLSIIGYIR",
		},
		{
			input:  Input6{s: "PAYPALISHIRING", numRows: 4},
			output: "PINALSIGYAHRPI",
		},
	}

	for _, test := range testCases {
		ans := ConvertZigZag(test.input.s, test.input.numRows)
		if ans != test.output {
			t.Errorf("expected %s got %s", test.output, ans)
		}
	}
}
