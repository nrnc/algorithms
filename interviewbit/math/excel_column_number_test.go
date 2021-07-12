package math

import "testing"

type ProblemExcelColumnNumber struct {
	input  string
	output int
}

func TestExcelColumNumber(t *testing.T) {
	testCases := []ProblemExcelColumnNumber{
		{
			input:  "C",
			output: 3,
		},
		{
			input:  "Z",
			output: 26,
		},
	}

	for _, test := range testCases {
		ans := ExcelColumNumber(test.input)
		if ans != test.output {
			t.Errorf("expected %d but got %d ", test.output, ans)
		}
	}
}
