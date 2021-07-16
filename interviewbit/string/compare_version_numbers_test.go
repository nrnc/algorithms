package string

import "testing"

type InputCompareVersion struct {
	A string
	B string
}

type ProblemCompareVersion struct {
	input  InputCompareVersion
	output int
}

func TestCompareVersion(t *testing.T) {
	testCases := []ProblemCompareVersion{
		{
			input: InputCompareVersion{
				A: "1",
				B: "01",
			},
			output: 0,
		},
		{
			input: InputCompareVersion{
				A: "1.1",
				B: "01",
			},
			output: 1,
		},
		{
			input: InputCompareVersion{
				A: "1",
				B: "01.1",
			},
			output: -1,
		},
	}
	for _, test := range testCases {
		ans := CompareVersion(test.input.A, test.input.B)
		if ans != test.output {
			t.Errorf("expected %d but got %d", test.output, ans)
		}
	}
}
