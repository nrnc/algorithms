package math

import "strconv"

func FizzBuzz(A int) []string {
	ans := make([]string, 0)

	for i := 1; i <= A; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans = append(ans, "FizzBuzz")
		} else if i%3 == 0 {
			ans = append(ans, "Fizz")
		} else if i%5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}

	return ans
}
