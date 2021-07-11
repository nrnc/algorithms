package math

import (
	"reflect"
	"testing"
)

type ProblemFizzBuzz struct {
	input  int
	output []string
}

func TestFizzBuzz(t *testing.T) {
	testCases := []ProblemFizzBuzz{
		{
			input:  15,
			output: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, test := range testCases {
		ans := FizzBuzz(test.input)

		if !reflect.DeepEqual(ans, test.output) {
			t.Errorf("expected %v but got %v", test.output, ans)
		}
	}
}
