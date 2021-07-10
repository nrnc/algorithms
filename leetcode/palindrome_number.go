package leetcode

func IsPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	tmp := x
	var reverse int
	for tmp != 0 {
		reverse = reverse*10 + tmp%10
		tmp = tmp / 10
	}

	return reverse == x
}
