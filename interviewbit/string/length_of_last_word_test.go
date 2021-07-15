package string

import "testing"

type ProblemLengthOfLastWord struct {
	input  string
	output int
}

func TestLengthOfLastWord(t *testing.T) {
	testCases := []ProblemLengthOfLastWord{
		{
			input:  "Hello World",
			output: 5,
		},
		{
			input:  "Hello World       ",
			output: 5,
		},
	}

	for _, test := range testCases {
		ans := LengthOfLastWord(test.input)
		if ans != test.output {
			t.Errorf("expected %d but got %d", test.output, ans)
		}
	}
}
